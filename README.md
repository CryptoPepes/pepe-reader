# CryptoPepe Reader

Consists of 2 parts:

*Reader service (`/datastoring`)*
 
- connects to an ethereum node
- starts datastore worker
- backfills existing pepes into the datastore
- watches for new pepes, and pushes them into the datastore

*Pepe (`/pepe`)*
- parses DNA data into a Pepe / PepeLook struct, encode-able to JSON
- fully parses DNA into pepe properties.

See **datastore** documentation (bottom README) to run the emulator.

## Running

```bash

go run --rpc="wss://ropsten.infura.io/ws" \
--token-address="... address ..." \
--sale-auction-address="..." \
--cozy-auction-address="..." \
main.go

```

### Binding generation

A PR with a critical bug-fix is open: [go-ethereum # 15676](https://github.com/ethereum/go-ethereum/pull/15676)

See smartcontracts repo README for further instructions.


# Deploy

Swap file creation (building takes a lot of memory)

```bash
cd /var
touch swap.img
chmod 600 swap.img

# 4 GB of swap
dd if=/dev/zero of=/var/swap.img bs=1024k count=4000
mkswap /var/swap.img
swapon /var/swap.img
free

# Make swap persistent
echo "/var/swap.img    none    swap    sw    0    0" >> /etc/fstab
```

```bash
# login
eval $(ssh-agent)
ssh-add <your key path>
ssh root@cryptopepes.io

# --- in remote machine ---

# Create reader app
dokku apps:create reader

# Access to volume with the IPC
# If archive node: 
# dokku docker-options:add reader deploy,run "-v ethereum:/root/.ethereum"

# Configure start
# Light-mode:
dokku config:set reader DOKKU_DOCKERFILE_START_CMD="--rpc=wss://ropsten.infura.io/ws \
--token-address=0x966383a597372cd2dea4d247a69db5a1fce8d3da \
--sale-auction-address=0xb1a4c9e7c5fb6866c1103a239f3cd333d36276d2 \
--cozy-auction-address=0x8fd285424995dd2adace2a3d0acb550717690367 \
--backfills=false"
# Full/backfill mode:
dokku config:set --no-restart reader DOKKU_DOCKERFILE_START_CMD="--rpc=/root/.ethereum/testnet/geth.ipc \
--token-address=0x966383a597372cd2dea4d247a69db5a1fce8d3da \
--sale-auction-address=0xb1a4c9e7c5fb6866c1103a239f3cd333d36276d2 \
--cozy-auction-address=0x8fd285424995dd2adace2a3d0acb550717690367 \
--backfills=true"

# main net
dokku config:set reader DOKKU_DOCKERFILE_START_CMD="--rpc=wss://ropsten.infura.io/ws \
--token-address=0x84ac94f17622241f313511b629e5e98f489ad6e4 \
--sale-auction-address=0x28ae3df366726d248c57b19fa36f6d9c228248be \
--cozy-auction-address=0xe2c43d2c6d6875c8f24855054d77b5664c7e810f \
--backfills=false"
# Full/backfill mode:
dokku config:set --no-restart reader DOKKU_DOCKERFILE_START_CMD="--rpc=/root/.ethereum/geth.ipc \
--token-address=0x84ac94f17622241f313511b629e5e98f489ad6e4 \
--sale-auction-address=0x28ae3df366726d248c57b19fa36f6d9c228248be \
--cozy-auction-address=0xe2c43d2c6d6875c8f24855054d77b5664c7e810f \
--backfills=true"

dokku config:set reader GOOGLE_APPLICATION_CREDENTIALS="datastore-key.json"
dokku config:set reader DATASTORE_PROJECT_ID="cryptopepe-192921"
dokku config:set reader APP_PATH="/app/"

# back to local machine
> exit

# Go to reader project dir
git remote add ocean-reader dokku@188.166.127.9:reader
git remote add ocean-reader dokku@188.166.115.121:reader

# Deploy
git push ocean-reader master
```

Other:

```bash
# Get app docker container ids
dokku ls
# Attach to running process, without hooking signals; this makes exiting easier & less error prone
docker attach --sig-proxy=false <container-id>
# Enter container (Using sh shell, as bash is not included)
docker exec -it <container-id> sh
```

## Notes / Gotchas

1) `govendor` is used to build dependencies
1) Go version for production is defined in `vendor/vendor.json/heroku.goVersion`
1) Special govendor thing (geth requires a C-only
 package to be included, which govendor ignores by default):
 `govendor add github.com/ethereum/go-ethereum/crypto/secp256k1/libsecp256k1/^`
 
# Datastore

The production datastore key (`datastore-key.json`) gives access to the 
Please use the emulator, don't change the production DB unnecessarily.

```
# Start emulator
./datastore-emulator.sh
```

## Environment variables:

```bash
# When set, program uses the emulator
DATASTORE_EMULATOR_HOST=localhost:8081

# Go-lang only: Project ID, also required for the emulator
DATASTORE_PROJECT_ID=cryptopepe-192921
# NodeJS version:
GCLOUD_PROJECT=cryptopepe-192921
```


### Production

Export `GOOGLE_APPLICATION_CREDENTIALS=datastore-key.json` to access the Google Cloud datastore from the application.

#### Extra

Creating indexes from the `index.yml`:

```bash
CLOUDSDK_CORE_PROJECT=cryptopepe-main GOOGLE_APPLICATION_CREDENTIALS=./datastore-key.json gcloud datastore create-indexes datastore-emu/WEB-INF/index.yaml
```


# Setting up the backfill node

This node setup is for backfilling,
 hence the need for a node of our own capable of reading large amounts of contract history.

The node is accessible from IPC (same UID necessary, run within both docker containers as root).

## Approach

1) Setup a regular Dokku VPS
  - With swap file, in case build/sync takes more than available memory
2) Attach a storage block for the full ethereum chaindata (> 100 GB)
3) Create a Dokku app for the ethereum node
4) Set-up container volume configuration
5) Deploy go-ethereum from source, to Dokku ethereum app
6) Deploy crypto-pepe reader to its own dokku app, sharing the same docker volume
7) Cryptopepe-reader communicates via IPC with ethereum container

## Commands

To setup testnode in full-sync mode
 (not in archive gc-mode, older state trie is useless, only need old logs):

### Setup remote:

```bash
# Do not forget to set up a swap file (See instructions in Deploy section)

# Create testnode app
dokku apps:create mainnode

# Disable proxy, it's not a webserver, just use the direct ports (Also, proxy is only for http(s))
dokku proxy:disable mainnode
# Enable the dokku app to bind to external interfaces
dokku network:set mainnode bind-all-interfaces true
# RPC over http/ws is not on by default, as it should be, so no need to do anything for those ports.

# Make volume
mkdir /mnt/my-storage/ethereum
# note: -o specifies volume-driver options (key=value)
docker volume create -d local -o type=none -o o=bind -o device=/mnt/my-storage/ethereum ethereum
# configure persistent shared volume for chaindata
dokku docker-options:add mainnode deploy,run "-v ethereum:/root/.ethereum"
# Configure the testnode
dokku config:set --no-restart mainnode DOKKU_DOCKERFILE_START_CMD="--syncmode=fast --cache=6000"
```

### Deploy from local to remote:

```
# Go to geth dir
cd "$GOPATH/src/github.com/ethereum/go-ethereum"

# Checkout a stable version of choice
git checkout v1.8.16

# Add dokku test remote
git remote add ocean-testnode dokku@123.100.123.100:testnode

# Start ssh-agent, add ssh key

# Deploy (targetting master branch, from whatever branch to deploy with)
git push ocean-testnode <my branch>:master
# E.g.: git push ocean-mainnode v1.8.16:master


```


