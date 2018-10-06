package main

import (
	"github.com/essentiaone/ess-atomic-swap/config"
	"github.com/essentiaone/ess-atomic-swap/server"
)

func main() {
	settings, _ := config.Init()
	if err := server.Init(settings.Server); err != nil {
		panic(err)
	}
}
