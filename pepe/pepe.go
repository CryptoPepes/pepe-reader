package pepe

import (
	"math/big"
	"github.com/ethereum/go-ethereum/common"
)

type Pepe struct {
	Master        common.Address
	Genotype      [2]*big.Int
	CanCozyAgain  uint64
	Generation    uint64
	Father        *big.Int
	Mother        *big.Int
	PepeName      [32]byte
	CoolDownIndex uint8
}
