package btc

// BitcoinUseCase is an interface for using BR of BTC atomic swap
type BitcoinUseCase interface {
	CheckTxStatus(string) bool
}
