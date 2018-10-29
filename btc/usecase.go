package btc

// BitcoinUseCase is an interface for using BR of BTC atomic swap
type BitcoinUseCase interface {
	СheckTxStatus(string) bool
}
