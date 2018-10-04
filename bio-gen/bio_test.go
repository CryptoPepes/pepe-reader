package bio_gen

import (
  "testing"
  "cryptopepe.io/cryptopepe-reader/pepe"
  "math/rand"
  "fmt"
  "math/big"
)

func TestBioGenerator(t *testing.T) {
  bioGen := new(BioGenerator)
  bioGen.Load()

  for i := int64(0); i < 1000; i++ {
    dna := getRandomPepeDna(i)
    fmt.Printf("----- Pepe #%03d -----\n", i)
    fmt.Println(bioGen.ConvertDnaToBio(dna))
  }
}


func getRandomPepeDna(pepeId int64) *pepe.PepeDNA {

  var pepeRngSrc = rand.NewSource(pepeId)
  var pepeRng = rand.New(pepeRngSrc)

  dna := pepe.PepeDNA{
    random256Big(pepeRng), random256Big(pepeRng),
  }

  return &dna
}

func random256Big(rng *rand.Rand) *big.Int {
  res := new (big.Int)
  bits256 := make([]byte, 32, 32)
  rng.Read(bits256)
  res.SetBytes(bits256)
  return res
}
