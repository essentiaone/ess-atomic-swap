package eth

import (
	"github.com/essentiaone/ess-atomic-swap/models"
)

// EthereumRepository is an interface for storing info
type EthereumRepository interface {
	SaveEthereumTx(ethTx *models.EthereumTransaction)
	GetEthereumTx() *models.EthereumTransaction
}
