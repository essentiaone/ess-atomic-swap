package usecase

import (
	"github.com/essentiaone/ess-atomic-swap/btc"
	"github.com/essentiaone/ess-atomic-swap/models"
	"github.com/pkg/errors"
)

const (
	getTxReceiptMethodName = "gettransaction"
	txSuccess              = 6
)

// BitcoinUsecase contains interface that have save/get btc tx functions
type BitcoinUsecase struct {
	repository btc.BitcoinRepository
	node       btc.BitcoinNode
}

// New is a func that init BitcoinUsecase struct
func New(btcRep btc.BitcoinRepository, btcNode btc.BitcoinNode) btc.BitcoinUseCase {
	return &BitcoinUsecase{
		repository: btcRep,
		node:       btcNode,
	}
}

// CheckTxStatus func for implementing BitcoinUseCase interface
func (btc *BitcoinUsecase) CheckTxStatus(tx string) (bool, error) {
	bitcoinTransaction := &models.BitcoinTransaction{}
	response, err := btc.node.ExecuteRequest(getTxReceiptMethodName, bitcoinTransaction, tx)

	if response != nil && err != nil { // Error from ethereum node
		return false, err
	} else if err != nil {
		return false, errors.Wrap(err, models.ErrorServerInternal.Error())
	}

	transactionInfo := response.(*models.BitcoinTransaction)

	return transactionInfo.Confirmations >= txSuccess, nil
}
