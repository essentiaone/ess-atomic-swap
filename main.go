package main

import (
	"fmt"

	btc "github.com/essentiaone/ess-atomic-swap/btc/usecase"
	"github.com/essentiaone/ess-atomic-swap/config"
	eth "github.com/essentiaone/ess-atomic-swap/eth/usecase"
	"github.com/essentiaone/ess-atomic-swap/server"
)

const (
	ethNodeURL = "https://mainnet.infura.io/Q6negVFqtmL30lFN6rK5"
	ethTxHash  = "0xa2ec48862367f5a641b6d565bf6a4fbee067c9cc4647d9293eb15c4f4e47d2d5"
	btcNodeURL = "http://104.248.42.188:8332"
	btcTxHash  = "0c3533db373b45a5d11fdea12fe8a91ed5845cb5872918add54879a214a8a082"
)

func main() {
	settings, _ := config.Init()

	test := eth.New(nil, ethNodeURL)
	result := test.CheckTxStatus(ethTxHash)
	fmt.Println(result)

	testBitcoin := btc.New(nil, btcNodeURL)
	resultBitcoin := testBitcoin.СheckTxStatus(btcTxHash)
	fmt.Println(resultBitcoin)

	if err := server.Init(settings.Server); err != nil {
		panic(err)
	}
}
