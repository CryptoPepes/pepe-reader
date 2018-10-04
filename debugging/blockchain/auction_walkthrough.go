package main

import (
	"log"
	"cryptopepe.io/cryptopepe-reader/simulated"
	"math/big"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	"cryptopepe.io/cryptopepe-reader/debugging/blockchain/util"
)

func main() {

	setup := simulated.NewSimulatedSetup()

	count := big.NewInt(3)
	premineTx, err := setup.PepeTokenOwnerSession.PepePremine(count)
	if err != nil {
		log.Fatalf("Error premine transaction! %v", err)
	}
	fmt.Println("Premine gas: ", premineTx.Gas())
	setup.Commit()


	util.PrintWalletOf(setup, setup.AuctionOwnerSession.TransactOpts.From)
	util.PrintWalletOf(setup, setup.Alice.From)

	pepe1 := big.NewInt(1)

	//approve that the contract can take the pepe
	setup.PepeTokenOwnerSession.Approve(setup.AuctionAddr, pepe1)
	//ask the contract to start an auction, taking the pepe which we approved.
	auctionTx, err := setup.AuctionOwnerSession.StartAuction(pepe1, big.NewInt(10 * params.Finney), big.NewInt(10 * params.Finney), big.NewInt(60*60))
	if err != nil {
		log.Fatalln("Error, could not start auction!", err)
	}
	fmt.Println("Start-auction gas: ", auctionTx.Gas())
	setup.Commit()

	buyOptions := bind.TransactOpts{
		From:     setup.Alice.From,
		Signer:   setup.Alice.Signer,
		Value: big.NewInt(10 * params.Finney),
	}

	buyTx, err := setup.Auction.BuyPepe(&buyOptions, pepe1)
	if err != nil {
		log.Fatalln("Error, could not buy pepe 1!", err)
	}
	fmt.Println("Buy pepe value: ", buyTx.Value())
	fmt.Println("Buy pepe gas: ", buyTx.Gas())
	setup.Commit()

	util.PrintWalletOf(setup, setup.AuctionOwnerSession.TransactOpts.From)
	util.PrintWalletOf(setup, setup.Alice.From)


	fmt.Println("end")

}
