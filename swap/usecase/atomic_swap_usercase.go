package usecase

import (
	"github.com/essentiaone/ess-atomic-swap/models"
	"github.com/essentiaone/ess-atomic-swap/swap"
)

// AtomicSwapUsecase consist all dependency for work
type AtomicSwapUsecase struct {
	repository swap.AtomicSwapRepository
}

// New create AtomicSwapUsecase
func New(swapRep swap.AtomicSwapRepository) swap.AtomicSwapUseCase {
	return &AtomicSwapUsecase{
		repository: swapRep,
	}
}

// Initiate accept order and save it temporary table
func (swap *AtomicSwapUsecase) Initiate(*models.AtomicSwapInitiate) (string, error) {
	rawHash, err := swap.repository.SaveOrderTemporary()
	if err != nil {
		return "", err
	}

	return string(rawHash), nil
}
