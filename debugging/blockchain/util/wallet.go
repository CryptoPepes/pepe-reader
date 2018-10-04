package util

import (
	"cryptopepe.io/cryptopepe-reader/simulated"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"fmt"
	"math/big"
)

func PrintWalletOf(setup *simulated.SimulatedSetup, owner common.Address) {
	balance, err := setup.PepeTokenOwnerSession.BalanceOf(owner)
	if err != nil {
		log.Fatalf("Failed retrieving balance! %v", err)
	}
	balanceInt := balance.Int64()
	fmt.Printf("Printing pepe wallet of %s (balance: %d): \n", owner.String(), balanceInt)
	for i := int64(0); i < balanceInt; i++ {
		pepeId, err := setup.PepeTokenOwnerSession.TokenOfOwnerByIndex(owner, big.NewInt(i))
		if err != nil {
			log.Fatalf("Failed retrieving pepe #%d from wallet! %v", i, err)
		}
		fmt.Println(" -> Pepe ", pepeId)
		pepeData, err := setup.PepeTokenOwnerSession.GetPepe(pepeId)
		if err != nil {
			log.Fatalf("Failed retrieving pepe data for pepe %d! %v", pepeId.Int64(), err)
		}
		fmt.Printf("    ---> With genotype: %v\n", pepeData.Genotype)
	}
}
