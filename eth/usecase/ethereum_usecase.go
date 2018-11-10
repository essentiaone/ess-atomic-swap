package usecase

import (
	"github.com/essentiaone/ess-atomic-swap/eth"
	"github.com/essentiaone/ess-atomic-swap/models"
	"github.com/pkg/errors"
)

const (
	getTxReceiptMethodName = "eth_getTransactionReceipt"
	txSuccess              = "0x1"
)

// EthereumUseCase consist all dependency for work
type EthereumUseCase struct {
	node eth.EthereumNode
}

// New create EthereumUseCase
func New(ethNode eth.EthereumNode) eth.EthereumUseCase {
	return &EthereumUseCase{
		node: ethNode,
	}
}

// TxStatus check success status for Ethereum transaction
// TODO implement sturct for response with status error and description
func (eth *EthereumUseCase) TxStatus(tx string) (bool, error) {
	ethereumTransaction := &models.EthereumTransaction{}
	response, err := eth.node.ExecuteRequest(getTxReceiptMethodName, ethereumTransaction, tx)

	if response != nil && err != nil { // Error from ethereum node
		return false, err
	} else if err != nil {
		return false, errors.Wrap(err, models.ErrorServerInternal.Error())
	}

	transactionInfo := response.(*models.EthereumTransaction)

	return transactionInfo.Status == txSuccess, nil
}
