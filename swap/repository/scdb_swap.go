package repository

import (
	"github.com/essentiaone/ess-atomic-swap/swap"
)

type SCDBInfrastructure interface {
	Save()
	Get()
	Delete()
	Update()
}

type scdbSwapRepository struct {
	scdb SCDBInfrastructure
}

func New(scdb SCDBInfrastructure) swap.AtomicSwapRepository {
	return &scdbSwapRepository{
		scdb: scdb,
	}
}

func (scdb *scdbSwapRepository) SaveOrderTemporary() ([]byte, error) {
	return []byte{}, nil
}

func (scdb *scdbSwapRepository) MoveOrderInvariable() ([]byte, error) {
	return []byte{}, nil
}
