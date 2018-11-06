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

	testEthereum := eth.New(nil, ethereumNode)
	resultEthereum, err := testEthereum.CheckTxStatus(ethTxHash)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultEthereum)

	testBitcoin := btc.New(nil, bitcoinNode)
	resultBitcoin, err := testBitcoin.CheckTxStatus(btcTxHash)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultBitcoin)

	if err := server.Init(config.Server); err != nil {
		panic(err)
	}
}
