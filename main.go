package main

import (
	"fmt"

	"github.com/essentiaone/ess-atomic-swap/config"
	"github.com/essentiaone/ess-atomic-swap/eth/usecase"
	"github.com/essentiaone/ess-atomic-swap/server"
)

func main() {
	settings, _ := config.Init()

	test := usecase.New(nil)
	result := test.CheckTxStatus("0xa2ec48862367f5a641b6d565bf6a4fbee067c9cc4647d9293eb15c4f4e47d2d5")
	fmt.Println(result)

	if err := server.Init(settings.Server); err != nil {
		panic(err)
	}
}
