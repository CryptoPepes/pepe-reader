package main

import "cryptopepe.io/cryptopepe-reader/simulated"

func main()  {

	setup := simulated.NewSimulatedSetup()

	setup.Premine()
	setup.TestBreeding()
	setup.DistributePepes()
	setup.TestAuctions()

}