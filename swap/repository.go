package swap

import "github.com/essentiaone/ess-atomic-swap/models"

// AtomicSwapRepository is an interface for storing info
type AtomicSwapRepository interface {
	SaveOrderTemporary(*models.AtomicSwapInitiate) (string, error)
}
