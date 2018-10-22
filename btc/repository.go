package btc

import (
	"github.com/essentiaone/ess-atomic-swap/models"
)

// BitcoinRepository is an interface for storing info
type BitcoinRepository interface {
	saveBTCTx(btcTx models.BitcoinTransaction)
	getBTCTx() models.BitcoinTransaction
}
