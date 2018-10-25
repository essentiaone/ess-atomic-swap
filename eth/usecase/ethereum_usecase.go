package usecase

import (
	"encoding/json"

	"github.com/essentiaone/ess-atomic-swap/common"
	"github.com/essentiaone/ess-atomic-swap/eth"
	"github.com/essentiaone/ess-atomic-swap/models"
)

const (
	txSuccess = "0x1"
)

// EthereumUseCase consist all dependency for business logic
type EthereumUseCase struct {
	NodeAddress        string
	EthereumRepository eth.EthereumRepository
}

// New create EthereumUseCase
func New(ethRep eth.EthereumRepository, nodeAddress string) eth.EthrerumUseCase {
	return &EthereumUseCase{
		EthereumRepository: ethRep,
		NodeAddress:        nodeAddress,
	}
}

// CheckTxStatus check success status for ethereum transaction
// TODO implement sturct for response with status error and description
func (eth *EthereumUseCase) CheckTxStatus(tx string) bool {
	response, err := getTransactionInfo(eth.NodeAddress, tx)
	if err != nil {
		return false
	}

	transactionInfo, ok := response.Result.(*models.EthereumTransaction)
	if !ok {
		return false
	}

	if transactionInfo.Status == txSuccess {
		return true
	}

	return false
}

func getTransactionInfo(address, tx string) (*common.ResponseEthereumNode, error) {
	dataRequest := &common.RequestEthereumBody{
		Params:  []string{tx},
		Address: address,
		JSONRPC: "2.0",
		Method:  "eth_getTransactionReceipt",
		ID:      1,
	}

	responseBody, err := common.HTTPRequestToNode(dataRequest)
	defer func() {
		if responseBody != nil {
			responseBody.Close()
		}
	}()

	if err != nil {
		return nil, err
	}

	ethResponse := &common.ResponseEthereumNode{
		Result: &models.EthereumTransaction{},
	}
	err = json.NewDecoder(responseBody).Decode(ethResponse)
	if err != nil {
		return nil, err
	}

	return ethResponse, nil
}
