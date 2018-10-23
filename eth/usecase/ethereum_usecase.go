package usecase

import (
	"encoding/json"

	"github.com/essentiaone/ess-atomic-swap/models"

	"github.com/essentiaone/ess-atomic-swap/common"
	"github.com/essentiaone/ess-atomic-swap/eth"
)

const (
	address   = "https://mainnet.infura.io/Q6negVFqtmL30lFN6rK5 "
	txSuccess = "0x1"
	txFailed  = "0x0"
)

// EthereumUseCase consist all dependency for business logic
type EthereumUseCase struct {
	EthereumRepository eth.EthereumRepository
}

// New create EthereumUseCase
func New(ethRep eth.EthereumRepository) eth.EthrerumUseCase {
	return &EthereumUseCase{
		EthereumRepository: ethRep,
	}
}

// CheckTxStatus check success status for ethereum transaction
func (eth *EthereumUseCase) CheckTxStatus(tx string) bool {
	response, _ := eth.getTransactionInfo(address, tx)
	transactionInfo, ok := response.Result.(*models.EthereumTransaction)
	if !ok {
		return false
	}

	if transactionInfo.Status == txSuccess {
		return true
	}

	return false
}

func (eth *EthereumUseCase) getTransactionInfo(address, tx string) (*common.ResponseEthereumNode, error) {
	dataRequest := &common.RequestEthereumBody{
		Params:  []string{tx},
		Address: address,
		JSONRPC: "2.0",
		Method:  "eth_getTransactionReceipt",
		ID:      1,
	}

	responseBody, err := common.HTTPRequestToNode(dataRequest)
	if err != nil {
		return nil, err
	}
	defer responseBody.Close()

	ethResponse := &common.ResponseEthereumNode{
		Result: &models.EthereumTransaction{},
	}
	err = json.NewDecoder(responseBody).Decode(ethResponse)
	if err != nil {
		return nil, err
	}

	return ethResponse, nil
}
