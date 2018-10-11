package main

import (
	"cryptopepe.io/cryptopepe-reader/reader"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"flag"
	"cryptopepe.io/cryptopepe-reader/simulated"
	"fmt"
	"cryptopepe.io/cryptopepe-reader/datastoring"
	"cryptopepe.io/cryptopepe-reader/abi/token"
	"io/ioutil"
	"path"
	"encoding/json"
	"cryptopepe.io/cryptopepe-reader/chainutil"
	"cryptopepe.io/cryptopepe-reader/abi/sale"
	"cryptopepe.io/cryptopepe-reader/abi/cozy"
	"github.com/ethereum/go-ethereum/rpc"
	"os"
)

func main() {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Go working dir: ", dir)

	gethSimMode := flag.Bool("gethsim",
		false,
		"Activates simulated backend (Geth based)," +
			" ignores the --token-address and --ipc options if true.")
	trufflePath := flag.String("truffle",
		"",
		"Build path of truffle. If set, reads settings from truffle build output.")
	networkId := flag.String("truffle-net",
		"5777",
		"Network id to read from when retrieving deploy data from truffle output.")
	rpcAddress := flag.String("rpc",
		"http://localhost:8545",
			"http/ws/ipc RPC connection to full node.")
	tokenAddress := flag.String("token-address",
		"0x1234abcd1234abcd1234abcd1234abcd1234abcd",
		"Pepe token address.")
	saleAuctionAddress := flag.String("sale-auction-address",
		"0x1234abcd1234abcd1234abcd1234abcd1234abcd",
		"Pepe sale auction token address.")
	cozyAuctionAddress := flag.String("cozy-auction-address",
		"0x1234abcd1234abcd1234abcd1234abcd1234abcd",
		"Cozy time auction token address.")

	contractBaseBlock := flag.Uint64("contract-base-block",
		6482000,
		"The minimum block, backfills are only done starting from here.")

	backfillMode := flag.Bool("backfills",
		true,
		"If backfills should be ran (on schedule)")

	//parse flags!
	flag.Parse()

	// parse truffle data if available.
	if *trufflePath != "" {
		*tokenAddress, *saleAuctionAddress, *cozyAuctionAddress = ReadAddrFromTruffle(*trufflePath, *networkId)
	}

	fmt.Printf(`Started!
		Geth sim. mode: %v
		truffle path: %v
		truffle net id: %v
		rpc: %s
		token address: %s
		saleAuctionAddress: %s
		cozyAuctionAddress: %s
		contract base block number: %d
`,
		*gethSimMode, *trufflePath, *networkId, *rpcAddress, *tokenAddress, *saleAuctionAddress, *cozyAuctionAddress, *contractBaseBlock)


	var rawPepeToken *token.Token
	var rawSaleAuctionToken *sale.Sale
	var rawCozyAuctionToken *cozy.Cozy
	var chainInfo chainutil.ChainInfo

	if *gethSimMode {
		setup := simulated.NewSimulatedSetup()

		chainInfo = setup

		setup.Premine()
		setup.TestBreeding()
		setup.DistributePepes()
		setup.TestAuctions()

		rawPepeToken = setup.PepeBase
		rawSaleAuctionToken = setup.SaleAuction
		rawCozyAuctionToken = setup.CozyAuction

	} else {
		rpcClient, ethClient := NewRpcConnection(*rpcAddress)

		chainInfo = reader.NewClientChainInfo(rpcClient)

		var err error
		// Instantiate the contract
		rawPepeToken, err = token.NewToken(common.HexToAddress(*tokenAddress), ethClient)
		if err != nil {
			log.Fatalf("Failed to instantiate a Token contract: %v", err)
		}
		rawSaleAuctionToken, err = sale.NewSale(common.HexToAddress(*saleAuctionAddress), ethClient)
		if err != nil {
			log.Fatalf("Failed to instantiate a Token contract: %v", err)
		}
		rawCozyAuctionToken, err = cozy.NewCozy(common.HexToAddress(*cozyAuctionAddress), ethClient)
		if err != nil {
			log.Fatalf("Failed to instantiate a Token contract: %v", err)
		}
	}

	r := reader.NewChainReader(rawPepeToken, rawSaleAuctionToken, rawCozyAuctionToken, chainInfo)

	dc := datastoring.NewDataStoreClient()

	worker := datastoring.NewPepeDataWorker(r, dc, *contractBaseBlock)
	worker.StartSchedule(*backfillMode)

	//TODO add dynamic input thing to close worker with? (worker.Close())
}

// Connect to a full node, and bind to it.
func NewRpcConnection(rpcAddress string) (*rpc.Client, *ethclient.Client) {

	c, err := rpc.Dial(rpcAddress)
	if err != nil {
		log.Fatalln("Could not connect to RPC!", err)
	}

	// Create a RPC connection to a full node
	conn := ethclient.NewClient(c)

	fmt.Println("Dial ok!")

	return c, conn
}

func ReadAddrFromTruffle(trufflePath string, networkId string) (tokenAddress string, saleAuctionAddress string, cozyAuctionAddress string) {
	tokenConf := ReadTruffleJson(path.Join(trufflePath, "PepeBase.json"))
	saleAuctionConf := ReadTruffleJson(path.Join(trufflePath, "PepeAuctionSale.json"))
	cozyAuctionConf := ReadTruffleJson(path.Join(trufflePath, "CozyTimeAuction.json"))
	tokenAddress = tokenConf["networks"].(map[string]interface{})[networkId].(map[string]interface{})["address"].(string)
	saleAuctionAddress = saleAuctionConf["networks"].(map[string]interface{})[networkId].(map[string]interface{})["address"].(string)
	cozyAuctionAddress = cozyAuctionConf["networks"].(map[string]interface{})[networkId].(map[string]interface{})["address"].(string)
	return tokenAddress, saleAuctionAddress, cozyAuctionAddress
}

func ReadTruffleJson(path string) (map[string]interface{}) {
	byt, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln("Failed to load token truffle file.", path, err)
	}
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		log.Fatalln("Failed to parse token truffle config data.", path, err)
	}
	return dat
}