package main

import (
	"fmt"

	"github.com/essentiaone/ess-atomic-swap/rpc"

	btc "github.com/essentiaone/ess-atomic-swap/btc/usecase"
	"github.com/essentiaone/ess-atomic-swap/config"
	eth "github.com/essentiaone/ess-atomic-swap/eth/usecase"
	"github.com/essentiaone/ess-atomic-swap/server"
)

const (
	ethTxHash = "0xa2ec48862367f5a641b6d565bf6a4fbee067c9cc4647d9293eb15c4f4e47d2d5"
	btcTxHash = "0c3533db373b45a5d11fdea12fe8a91ed5845cb5872918add54879a214a8a082"
)

func main() {
	config, _ := config.Init()

	ethereumNode := rpc.New(config.EthereumNodeAddr)
	bitcoinNode := rpc.New(config.BitcoinNodeAddr)

	test := eth.New(nil, ethereumNode)
	result := test.CheckTxStatus(ethTxHash)
	fmt.Println(result)

	testBitcoin := btc.New(nil, bitcoinNode)
	resultBitcoin := testBitcoin.CheckTxStatus(btcTxHash)
	fmt.Println(resultBitcoin)

	if err := server.Init(config.Server); err != nil {
		panic(err)
	}
}
