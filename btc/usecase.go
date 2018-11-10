package btc

// BitcoinUseCase interface for work with Bitcoin blockchain
type BitcoinUseCase interface {
	TxStatus(string) (bool, error)
}
