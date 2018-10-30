package usecase

import (
	"encoding/json"

	"github.com/essentiaone/ess-atomic-swap/btc"
	"github.com/essentiaone/ess-atomic-swap/common"
	"github.com/essentiaone/ess-atomic-swap/models"
)

const (
	getTxReceiptMethodName = "gettransaction"
	txSuccess              = 4
)

// BitcoinUsecase contains interface that have save/get btc tx functions
type BitcoinUsecase struct {
	NodeAddress       string
	BitcoinRepository btc.BitcoinRepository
}

// New is a func that init BitcoinUsecase struct
func New(btcRep btc.BitcoinRepository, nodeAddress string) btc.BitcoinUseCase {
	return &BitcoinUsecase{
		BitcoinRepository: btcRep,
		NodeAddress:       nodeAddress,
	}
}

// СheckTxStatus func for implementing BitcoinUseCase interface
func (b *BitcoinUsecase) СheckTxStatus(tx string) bool {
	response, err := getTransactionInfo(b.NodeAddress, tx)
	if err != nil {
		return false
	}

	transactionInfo, ok := response.Result.(*models.BitcoinTransaction)
	if !ok {
		return false
	}

	if transactionInfo.Confirmations >= txSuccess {
		return true
	}

	return false
}

func getTransactionInfo(address, tx string) (*common.ResponseBitcoinNode, error) {
	dataRequest := &common.RequestBitcoinBody{
		Params:  []string{tx},
		Address: address,
		JSONRPC: "2.0",
		Method:  getTxReceiptMethodName,
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

	btcResponse := &common.ResponseBitcoinNode{
		Result: &models.BitcoinTransaction{},
	}
	err = json.NewDecoder(responseBody).Decode(btcResponse)
	if err != nil {
		return nil, err
	}

	return btcResponse, nil
}
