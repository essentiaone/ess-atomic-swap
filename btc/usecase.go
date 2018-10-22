package btc

import (
	"github.com/essentiaone/ess-atomic-swap/models"
)

type BTCUseCase interface {
	checkStatus(models.BTCTransaction)
}
