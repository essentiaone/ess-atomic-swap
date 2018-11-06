package swap

import (
	"github.com/essentiaone/ess-atomic-swap/models"
)

// AtomicSwapUseCase main logic for create atomic-swap
type AtomicSwapUseCase interface {
	Initiate(*models.AtomicSwapInitiate) (string, error)
}
