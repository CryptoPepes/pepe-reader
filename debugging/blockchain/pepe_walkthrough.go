package main

import (
	"log"
	"cryptopepe.io/cryptopepe-reader/simulated"
	"math/big"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
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
	setup.Commit()
	setup.Commit()

	motherId := big.NewInt(1)
	fatherId := big.NewInt(2)
	tx, err := setup.PepeTokenOwnerSession.CozyTime(motherId, fatherId, setup.PepeTokenOwnerSession.TransactOpts.From)
	if err != nil {
		log.Fatalf("Error cozy time transaction! %v", err)
	}
	setup.Commit()
	fmt.Println("Cozy time gas: ", tx.Gas())

	/**
		Breeding debugging
	 */

	motherPepe, err := setup.PepeTokenOwnerSession.GetPepe(big.NewInt(1))
	if err != nil {
		log.Fatalf("Error failed pepe #1! %v", err)
	}
	fmt.Printf("Mother genotype: %v\n", motherPepe.Genotype)

	fatherPepe, err := setup.PepeTokenOwnerSession.GetPepe(big.NewInt(2))
	if err != nil {
		log.Fatalf("Error failed pepe #2! %v", err)
	}
	fmt.Printf("Father genotype: %v\n", fatherPepe.Genotype)


	childPepe, err := setup.PepeTokenOwnerSession.GetPepe(big.NewInt(3))
	if err != nil {
		log.Fatalf("Error failed pepe #3! %v", err)
	}
	fmt.Printf("Child genotype: %v\n", childPepe.Genotype)

	printDnaAnalysis(motherPepe.Genotype, fatherPepe.Genotype, childPepe.Genotype)


	/**
	     Wallet debugging
	 */
	testPersonKey, _ := crypto.GenerateKey()
	testPerson := bind.NewKeyedTransactor(testPersonKey)

	fmt.Println("-------- Before transaction -----------")
	util.PrintWalletOf(setup, setup.ContractAuth.From)
	util.PrintWalletOf(setup, testPerson.From)

	fmt.Println("-------- transaction -----------")
	pepeTx, err := setup.PepeTokenOwnerSession.Transfer(testPerson.From, big.NewInt(1))
	if err != nil {
		log.Fatalf("Failed pepe transfer! %v", err)
	}
	fmt.Println(pepeTx)

	//Mine the transaction
	setup.Commit()

	fmt.Println("-------- After transaction -----------")
	util.PrintWalletOf(setup, setup.ContractAuth.From)
	util.PrintWalletOf(setup, testPerson.From)

	// Transfer a pepe to another new address


	fmt.Println("end")

}


func printDnaAnalysis(mother [2]*big.Int, father [2]*big.Int, child [2]*big.Int) {

	motherStr0 := ""
	motherStr1 := ""
	fatherStr0 := ""
	fatherStr1 := ""
	childStr0 := ""
	childStr1 := ""

	child0XorMother0 := ""
	child0XorMother1 := ""
	child1XorMother0 := ""
	child1XorMother1 := ""

	child0XorFather0 := ""
	child0XorFather1 := ""
	child1XorFather0 := ""
	child1XorFather1 := ""

	motherStr0 += fmt.Sprintf("%064x ", mother[0])
	motherStr1 += fmt.Sprintf("%064x ", mother[1])
	fatherStr0 += fmt.Sprintf("%064x ", father[0])
	fatherStr1 += fmt.Sprintf("%064x ", father[1])
	childStr0 += fmt.Sprintf("%064x ", child[0])
	childStr1 += fmt.Sprintf("%064x ", child[1])

	child0XorMother0 += fmt.Sprintf("%064x ", new(big.Int).Xor(child[0], mother[0]))
	child0XorMother1 += fmt.Sprintf("%064x ", new(big.Int).Xor(child[0], mother[1]))
	child1XorMother0 += fmt.Sprintf("%064x ", new(big.Int).Xor(child[1], mother[0]))
	child1XorMother1 += fmt.Sprintf("%064x ", new(big.Int).Xor(child[1], mother[1]))

	child0XorFather0 += fmt.Sprintf("%064x ", new(big.Int).Xor(child[0], father[0]))
	child0XorFather1 += fmt.Sprintf("%064x ", new(big.Int).Xor(child[0], father[1]))
	child1XorFather0 += fmt.Sprintf("%064x ", new(big.Int).Xor(child[1], father[0]))
	child1XorFather1 += fmt.Sprintf("%064x ", new(big.Int).Xor(child[1], father[1]))

	motherStr0 +=  "  "
	motherStr1 +=  "  "
	fatherStr0 +=  "  "
	fatherStr1 +=  "  "
	childStr0 +=  "  "
	childStr1 +=  "  "
	child0XorMother0 += "  "
	child0XorMother1 += "  "
	child1XorMother0 += "  "
	child1XorMother1 += "  "
	child0XorFather0 += "  "
	child0XorFather1 += "  "
	child1XorFather0 += "  "
	child1XorFather1 += "  "

	motherStr0 += fmt.Sprintf("%0256b ", mother[0])
	motherStr1 += fmt.Sprintf("%0256b ", mother[1])
	fatherStr0 += fmt.Sprintf("%0256b ", father[0])
	fatherStr1 += fmt.Sprintf("%0256b ", father[1])
	childStr0 += fmt.Sprintf("%0256b ", child[0])
	childStr1 += fmt.Sprintf("%0256b ", child[1])

	child0XorMother0 += fmt.Sprintf("%0256b ", new(big.Int).Xor(child[0], mother[0]))
	child0XorMother1 += fmt.Sprintf("%0256b ", new(big.Int).Xor(child[0], mother[1]))
	child1XorMother0 += fmt.Sprintf("%0256b ", new(big.Int).Xor(child[1], mother[0]))
	child1XorMother1 += fmt.Sprintf("%0256b ", new(big.Int).Xor(child[1], mother[1]))

	child0XorFather0 += fmt.Sprintf("%0256b ", new(big.Int).Xor(child[0], father[0]))
	child0XorFather1 += fmt.Sprintf("%0256b ", new(big.Int).Xor(child[0], father[1]))
	child1XorFather0 += fmt.Sprintf("%0256b ", new(big.Int).Xor(child[1], father[0]))
	child1XorFather1 += fmt.Sprintf("%0256b ", new(big.Int).Xor(child[1], father[1]))
	
	fmt.Println("mother 0:           " + motherStr0)
	fmt.Println("mother 1:           " + motherStr1)
	fmt.Println("father 0:           " + fatherStr0)
	fmt.Println("father 1:           " + fatherStr1)
	fmt.Println("child  0:           " + childStr0)
	fmt.Println("child  1:           " + childStr1)
	fmt.Println("----")
	fmt.Println("child 0 ^ mother 0: " + child0XorMother0)
	fmt.Println("child 0 ^ mother 1: " + child0XorMother1)
	fmt.Println("child 1 ^ mother 0: " + child1XorMother0)
	fmt.Println("child 1 ^ mother 0: " + child1XorMother1)
	fmt.Println("----")
	fmt.Println("child 0 ^ father 0: " + child0XorFather0)
	fmt.Println("child 0 ^ father 1: " + child0XorFather1)
	fmt.Println("child 1 ^ father 0: " + child1XorFather0)
	fmt.Println("child 1 ^ father 0: " + child1XorFather1)
}