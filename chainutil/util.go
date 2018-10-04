package chainutil

type ChainInfo interface {
	GetCurrentBlock() (uint64, error)
}
