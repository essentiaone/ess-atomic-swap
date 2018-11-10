package eth

// EthereumUseCase interface for work with Ethereum blockchain
type EthereumUseCase interface {
	TxStatus(string) (bool, error)
}
