package btc

import (
	"github.com/essentiaone/ess-atomic-swap/models"
)

type BTCRepository interface {
	saveBTCTx(btcTx models.BTCTransaction)
	getBTCTx() models.BTCTransaction
}
