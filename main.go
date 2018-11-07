package main

import (
	"github.com/essentiaone/ess-atomic-swap/config"
	"github.com/essentiaone/ess-atomic-swap/server"
)

func main() {
	config, err := config.Init()
	if err != nil {
		panic(err)
	}

	if err := server.Init(config.Server); err != nil {
		panic(err)
	}
}
