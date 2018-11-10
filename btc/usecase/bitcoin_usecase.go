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

// BitcoinUsecase consist all dependency for work
type BitcoinUsecase struct {
	node btc.BitcoinNode
}

// New is a func that init BitcoinUsecase struct
func New(btcNode btc.BitcoinNode) btc.BitcoinUseCase {
	return &BitcoinUsecase{
		node: btcNode,
	}
}

// TxStatus func for implementing BitcoinUseCase interface
func (btc *BitcoinUsecase) TxStatus(tx string) (bool, error) {
	bitcoinTransaction := &models.BitcoinTransaction{}
	response, err := btc.node.ExecuteRequest(getTxReceiptMethodName, bitcoinTransaction, tx)

	if response != nil && err != nil { // Error from Ethereum node
		return false, err
	} else if err != nil {
		return false, errors.Wrap(err, models.ErrorServerInternal.Error())
	}

	transactionInfo := response.(*models.BitcoinTransaction)

	return transactionInfo.Confirmations >= txSuccess, nil
}
