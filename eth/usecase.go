package eth

// EthrerumUseCase is an interface for using BR of ETH atomic swap
type EthrerumUseCase interface {
	CheckTxStatus(string) bool
}
