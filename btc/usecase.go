package btc

import (
	"github.com/essentiaone/ess-atomic-swap/models"
)

// BitcoinUseCase is an interface for using BR of BTC atomic swap
type BitcoinUseCase interface {
	checkStatus(models.BitcoinTransaction)
}
