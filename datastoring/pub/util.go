package pub

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"time"
	"context"
)

func FilterOpts(fromBlock uint64, toBlock uint64) *bind.FilterOpts {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	opts := bind.FilterOpts{
		Start: fromBlock,
		End: &toBlock,
		Context: ctx,
	}

	return &opts
}