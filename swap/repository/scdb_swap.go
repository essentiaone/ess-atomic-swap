package repository

import (
	"github.com/essentiaone/ess-atomic-swap/models"
	"github.com/essentiaone/ess-atomic-swap/swap"
)

type SCDBInfrastructure interface {
	SaveOrder(string, string, string, int64) (string, error)
}

type scdbSwapRepository struct {
	scdb SCDBInfrastructure
}

func New(scdb SCDBInfrastructure) swap.AtomicSwapRepository {
	return &scdbSwapRepository{
		scdb: scdb,
	}
}

func (db *scdbSwapRepository) SaveOrderTemporary(order *models.AtomicSwapInitiate) (string, error) {
	return db.scdb.SaveOrder(order.From, order.To, order.Password, int64(order.Amount))
}
