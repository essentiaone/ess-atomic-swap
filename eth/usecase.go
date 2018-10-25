package eth

// EthereumUseCase is an interface for using BR of ETH atomic swap
type EthereumUseCase interface {
	CheckTxStatus(string) bool
}
