package simulated


import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"context"
	"cryptopepe.io/cryptopepe-reader/abi/token"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"time"
	"math/rand"
	"errors"
	"cryptopepe.io/cryptopepe-reader/abi/sale"
	"cryptopepe.io/cryptopepe-reader/abi/cozy"
)

const (
	// Excl. contract owners, genesis auth, and Alice&Bob.
	UserCount = 20
	BreedRounds = 3
	BreedRoundPhases = 5
	BreedMaxPerPhase = 100

	TestSaleAuctionsElapsed = 30
	TestSaleAuctionsCancel = 10
	TestSaleAuctionsSold = 30

	TestCozyAuctionsElapsed = 30
	TestCozyAuctionsCancel = 10
	TestCozyAuctionsSold = 30
)

var rng = rand.New(rand.NewSource(1234))

type SimulatedSetup struct {
	PepeBase              *token.Token
	PepeBaseAddr          common.Address
	SaleAuction           *sale.Sale
	SaleAuctionAddr       common.Address
	CozyAuction           *cozy.Cozy
	CozyAuctionAddr       common.Address

	sim                   *backends.SimulatedBackend

	GenesisAuth           *bind.TransactOpts
	ContractAuth          *bind.TransactOpts
	Alice                 *bind.TransactOpts
	Bob                   *bind.TransactOpts
	// More test users, each allocated with 1000 Eth
	TestUsers             []*bind.TransactOpts

	PepeTokenOwnerSession *token.TokenSession
	SaleAuctionOwnerSession   *sale.SaleSession
	CozyAuctionOwnerSession   *cozy.CozySession


	blockNum              uint64
}

func NewSimulatedSetup() *SimulatedSetup {

	setup := new(SimulatedSetup)

	// Generate a new random account and a funded simulator
	setup.GenesisAuth = makeUserAuth()
	setup.ContractAuth = makeUserAuth()
	setup.Alice = makeUserAuth()
	setup.Bob = makeUserAuth()
	//setup.ContractAuth.GasLimit = big.NewInt()

	eth1000 := new(big.Int).Mul(big.NewInt(1000), big.NewInt(params.Ether))
	genesisAlloc := core.GenesisAlloc{
		setup.GenesisAuth.From: core.GenesisAccount{Balance: eth1000},
		setup.ContractAuth.From: core.GenesisAccount{Balance: eth1000},
		setup.Alice.From: core.GenesisAccount{Balance: eth1000},
		setup.Bob.From: core.GenesisAccount{Balance: eth1000},
	}
	userCount := UserCount
	setup.TestUsers = make([]*bind.TransactOpts, userCount, userCount)
	// Create #userCount more test-"users"
	for i := 0; i < userCount; i++ {
		user := makeUserAuth()
		setup.TestUsers[i] = user
		genesisAlloc[user.From] = core.GenesisAccount{Balance: eth1000}
	}
	// Create the backend, which creates an internal simulated blockchain, using our genesis alloc.
	setup.sim = backends.NewSimulatedBackend(genesisAlloc)

	nullAddr := common.HexToAddress("0x0")

	// Deploy a token contract on the simulated blockchain
	pepeTokenAddr, tx, pepeToken, err := token.DeployToken(setup.ContractAuth, setup.sim)
	if err != nil {
		log.Fatalf("Failed to deploy new PepeBase: %v", err)
	}
	setup.PepeBase = pepeToken
	setup.PepeBaseAddr = pepeTokenAddr
	fmt.Println("Deployed pepe token to: "+pepeTokenAddr.String())
	fmt.Println("Deploy tx: "+tx.String())

	setup.Commit()

	// Deploy auction contracts
	// -------
	// Pepe sale contract
	saleAuctionAddr, tx, saleAuctionContract, err := sale.DeploySale(setup.ContractAuth, setup.sim, pepeTokenAddr, nullAddr)
	if err != nil {
		log.Fatalf("Failed to deploy new PepeBase: %v", err)
	}
	setup.SaleAuction = saleAuctionContract
	setup.SaleAuctionAddr = saleAuctionAddr

	fmt.Println("Deployed sale auction contract to: "+saleAuctionAddr.String())
	fmt.Println("Deploy tx: "+tx.String())

	setup.Commit()

	// Pepe cozy contract
	cozyAuctionAddr, tx, cozyAuctionContract, err := cozy.DeployCozy(setup.ContractAuth, setup.sim, pepeTokenAddr, nullAddr)
	if err != nil {
		log.Fatalf("Failed to deploy new PepeBase: %v", err)
	}
	setup.CozyAuction = cozyAuctionContract
	setup.CozyAuctionAddr = cozyAuctionAddr

	fmt.Println("Deployed cozy auction contract to: "+cozyAuctionAddr.String())
	fmt.Println("Deploy tx: "+tx.String())

	setup.Commit()

	setup.PepeTokenOwnerSession = &token.TokenSession{
		Contract: setup.PepeBase,
		CallOpts: bind.CallOpts{
			Pending: false,
			Context: context.Background(),
		},
		TransactOpts: bind.TransactOpts{
			From:     setup.ContractAuth.From,
			Signer:   setup.ContractAuth.Signer,
		},
	}

	setup.SaleAuctionOwnerSession = &sale.SaleSession{
		Contract: setup.SaleAuction,
		CallOpts: bind.CallOpts{
			Pending: false,
			Context: context.Background(),
		},
		TransactOpts: bind.TransactOpts{
			From:     setup.ContractAuth.From,
			Signer:   setup.ContractAuth.Signer,
		},
	}

	setup.CozyAuctionOwnerSession = &cozy.CozySession{
		Contract: setup.CozyAuction,
		CallOpts: bind.CallOpts{
			Pending: false,
			Context: context.Background(),
		},
		TransactOpts: bind.TransactOpts{
			From:     setup.ContractAuth.From,
			Signer:   setup.ContractAuth.Signer,
		},
	}

	return setup
}

func (setup *SimulatedSetup) Commit() {
	setup.sim.Commit()
	setup.blockNum++
}

func (setup *SimulatedSetup) GetCurrentBlock() (uint64, error) {
	return setup.blockNum, nil
}

func makeUserAuth() *bind.TransactOpts {
	contractKey, _ := crypto.GenerateKey()
	return bind.NewKeyedTransactor(contractKey)
}

func (setup *SimulatedSetup) Premine() {
	balance, _ := setup.sim.BalanceAt(nil, setup.ContractAuth.From, nil)
	fmt.Println("Balance before premine: " + balance.String())
	for i := 0; i < 10; i++ {
		fmt.Printf("mining block for %d for pre-mine\n", i)
		premineAmount := big.NewInt(10)
		tx, err := setup.PepeTokenOwnerSession.PepePremine(premineAmount)
		if err != nil {
			log.Fatalf("Error pre-mining transaction! %v", err)
		}
		fmt.Println("Premine gas: ", tx.Gas())
		//mine the block
		setup.Commit()
	}
	balance, _ = setup.sim.BalanceAt(nil, setup.ContractAuth.From, nil)
	fmt.Println("Balance after premine: " + balance.String())
}

func (setup *SimulatedSetup) getPepeSupply() int64 {
	supply, err := setup.PepeTokenOwnerSession.TotalSupply()
	if err != nil {
		log.Fatalf("Error retrieving total supply err: %v", err)
	}
	return supply.Int64()
}


func (setup *SimulatedSetup) TestBreeding() {
	balance, _ := setup.sim.BalanceAt(nil, setup.ContractAuth.From, nil)
	fmt.Println("Balance before test breeding: " + balance.String())

	// Make free transactions, to enable us to generate a very large test set.
	freeTxOpts := &bind.TransactOpts{
		From: setup.ContractAuth.From,
		Signer:   setup.ContractAuth.Signer,
		GasPrice: big.NewInt(0),
		// Use a large gas limit, "costs" more,
		//  but prevents estimating gas by computing each transaction 10+ times.
		GasLimit: 600000,
	}

	// Repeat the whole process a few times, to have young low-gens as well.
	for round := 0; round < BreedRounds; round++ {
		for p := 0; p < BreedRoundPhases; p++ {

			//randomly breed one half of the pepes with the other half.
			totalPepes := setup.getPepeSupply()

			shuffled := rng.Perm(int(totalPepes))

			log.Printf("Round %d, phase %d, pepe supply: %d\n", round, p, totalPepes)

			// at most BreedMaxPerPhase (2 * BreedMaxPerPhase parents) random pepes per phase
			for i := 0; i <= 2 * BreedMaxPerPhase && i + 1 < len(shuffled); {
				mother := big.NewInt(int64(shuffled[i]))
				i++
				father := big.NewInt(int64(shuffled[i]))
				i++
				//mo, _ := setup.PepeTokenOwnerSession.GetPepe(mother)
				//fa, _ := setup.PepeTokenOwnerSession.GetPepe(father)

				_, err := setup.PepeBase.CozyTime(freeTxOpts, mother, father, freeTxOpts.From)
				if err != nil {
					log.Fatalf("Error breeding transaction! mother: %s, father: %s, err: %v", mother.Text(10), father.Text(10), err)
					//log.Println(mo, fa)
					//continue
				}

				// Every now and then, mine the pepes, and add some time.
				// Not every pepe of a generation is born in the same block / time.
				// Also, gas limits apply, stay below the limits.
				if i % 10 == 0 {
					// just skip by max cooldown.
					setup.Commit()
					setup.sim.AdjustTime(time.Hour * 24 * 7 + time.Second)
					setup.Commit()
				}
			}
			// Note mine the block first, then adjust time, then mine again.
			// This works-around losing the time adjustment on the pending block when registering a new transaction (possibly a bug in Geth?)
			setup.Commit()
			// skip ahead by max cooldown.
			setup.sim.AdjustTime(time.Hour * 24 * 7 + time.Second)
			//mine a block, gas is free, but born pepes txs need to be confirmed.
			setup.Commit()
		}
	}

}

func (setup *SimulatedSetup) DistributePepes() {
	userCount := int64(len(setup.TestUsers))
	reservedPepes := int64(10)
	totalPepes := setup.getPepeSupply() - reservedPepes
	pepePerUser := totalPepes / userCount
	if pepePerUser == 0 {
		log.Fatalln("Not enough pepes to distribute to test users!")
	}

	freeTxOpts := &bind.TransactOpts{
		From: setup.ContractAuth.From,
		Signer:   setup.ContractAuth.Signer,
		GasPrice: big.NewInt(0),
		GasLimit: 500000,// give more than enough gas, avoid computing gas costs, speed up.
	}

	k := reservedPepes
	for i := int64(0); i < userCount; i++ {
		user := setup.TestUsers[i]
		for j := int64(0); j < pepePerUser; j++ {
			k++
			// Mine the last few txs, stay below gas limit
			if k % 5 == 0 {
				setup.Commit()
			}
			setup.PepeBase.Transfer(freeTxOpts, user.From, big.NewInt(k))
		}
		log.Printf("Send %d pepes to user %d (addr: %s)", pepePerUser, i, user.From.String())
	}
	// finalize last
	setup.Commit()
}


func (setup *SimulatedSetup) TestAuctions() {
	log.Println("Starting test auctions!!!")

	// saleOrCozy: false is cozy
	setup.runAuctions(true,
		TestSaleAuctionsElapsed,
		TestSaleAuctionsCancel,
		TestSaleAuctionsSold)

	setup.runAuctions(false,
		TestCozyAuctionsElapsed,
		TestCozyAuctionsCancel,
		TestCozyAuctionsSold)

}

func (setup *SimulatedSetup) runAuctions(saleOrCozy bool, elapsedMax int, cancelMax int, soldMax int) {

	auctionedPepes := make(map[*big.Int]bool)

	var createTxOpts = func(maker int) *bind.TransactOpts {
		txOpts := &bind.TransactOpts{
			From:     setup.TestUsers[maker].From,
			Signer:   setup.TestUsers[maker].Signer,
			GasPrice: big.NewInt(0),
			GasLimit: 500000, // give more than enough gas, avoid computing gas costs, speed up.
		}
		return txOpts
	}

	var getPepeFromWallet = func(wallet common.Address) (*big.Int, error) {
		// Get balance in pepes
		pepeBalance, err := setup.PepeTokenOwnerSession.BalanceOf(wallet)
		if err != nil {
			return nil, err
		}
		pepeMax := pepeBalance.Int64()
		// Get random Pepe from wallet
		pepeId, err := setup.PepeTokenOwnerSession.TokenOfOwnerByIndex(wallet, big.NewInt(rng.Int63n(pepeMax)))
		if err != nil {
			return nil, err
		}
		return pepeId, nil
	}

	var startAuction = func(maker int) (id *big.Int, err error) {

		txOpts := createTxOpts(maker)

		pepeId, err := getPepeFromWallet(txOpts.From)
		if err != nil {
			return nil, err
		}

		if auctionedPepes[pepeId] {
			return nil, errors.New("pepe "+pepeId.String()+" is already being auctioned")
		}

		auctionedPepes[pepeId] = true

		approvAddr := setup.CozyAuctionAddr
		if saleOrCozy {
			approvAddr = setup.SaleAuctionAddr
		}
		// approve the auction contract to move the pepe
		if _, err := setup.PepeBase.Approve(txOpts, approvAddr, pepeId); err != nil {
			return nil, err
		}

		// mine approval tx
		setup.Commit()

		// start at 1 - 5 Eth
		startPrice := new(big.Int).Mul(big.NewInt(rng.Int63n(4000) + 1000), big.NewInt(params.Finney))
		// end at 0.1 - 7 Eth
		endPrice := new(big.Int).Mul(big.NewInt(rng.Int63n(5900) + 100), big.NewInt(params.Finney))

		if saleOrCozy {
			setup.SaleAuction.StartAuction(txOpts, pepeId, startPrice, endPrice, 3 * 60 * 60)
		} else {
			setup.CozyAuction.StartAuction(txOpts, pepeId, startPrice, endPrice, 3 * 60 * 60)
		}

		// Mine start-auction tx
		setup.Commit()

		return pepeId, nil
	}

	totalAuctions := elapsedMax + cancelMax + soldMax
	elapsed := 0
	cancel := 0
	sold := 0

	users := len(setup.TestUsers)

	// keep trying auctions, up to 2 times
	tries := totalAuctions * 2
	mainLoop: for i := 0; i < tries; i++ {

		maker := rng.Intn(users)

		pepeId, err := startAuction(maker)
		if err != nil {
			log.Printf("Could not start auction for pepe: %s, maker: %d", pepeId.String(), maker)
			continue
		}

		switch rng.Intn(3) {
		case 0: // elapse
			if elapsed < elapsedMax {
				elapsed++
				log.Printf("Started auction [sim. elapse], pepe: %s, user: %d.\n", pepeId, maker)
				break
			}
			fallthrough
		case 1: // cancel
			if cancel < cancelMax {
				cancel++
				txOpts := createTxOpts(maker)
				if saleOrCozy {
					setup.SaleAuction.SavePepe(txOpts, pepeId)
				} else {
					setup.CozyAuction.SavePepe(txOpts, pepeId)
				}
				// commit the cancel.
				setup.Commit()
				log.Printf("Started and cancelled auction, pepe: %s, user: %d.\n", pepeId, maker)
				break
			}
			fallthrough
		case 2: // sold
			if sold < soldMax {
				sold++
				taker := rng.Intn(users)
				txOpts := createTxOpts(taker)
				if saleOrCozy {
					if _, err := setup.SaleAuction.BuyPepe(txOpts, pepeId); err != nil {
						log.Printf("Could not buy pepe %s, taker: %d, maker: %d\n", pepeId.String(), taker, maker)
						break
					}
				} else {
					candidate, err := getPepeFromWallet(txOpts.From)
					if err != nil {
						log.Println("Could not get pepe candidate for cozy time from wallet! taker wallet: ", taker)
						break
					}
					if _, err := setup.CozyAuction.BuyCozy(txOpts, pepeId, candidate, true, txOpts.From); err != nil {
						log.Printf("Could not buy pepe %s, taker: %d, maker: %d\n", pepeId.String(), taker, maker)
						break
					}
				}
				// commit the auction.
				setup.Commit()
				log.Printf("Started and finalized auction, pepe: %s, maker: %d, taker: %d.\n", pepeId, maker, taker)
				break
			}
			fallthrough
		default:
			// no action types left.
			break mainLoop
		}
	}
}
