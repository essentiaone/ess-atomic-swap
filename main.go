package main

import (
	"fmt"

	"github.com/essentiaone/ess-atomic-swap/common"
	"github.com/essentiaone/ess-atomic-swap/config"
	"github.com/essentiaone/ess-atomic-swap/logger"
	"github.com/essentiaone/ess-atomic-swap/server"
	"github.com/essentiaone/ess-atomic-swap/swap/infrastructure"
	"github.com/essentiaone/ess-atomic-swap/swap/repository"
	"github.com/essentiaone/ess-atomic-swap/swap/usecase"
)

func main() {
	config, err := config.Init()
	if err != nil {
		panic(err)
	}

	logger, err := logger.New(config.IsProduction)
	if err != nil {
		panic(err)
	}

	dbInfra, err := infrastructure.New(
		config.EthereumNodeAddr,
		config.SCDBAddr,
		config.PrivateKet,
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	db := repository.New(dbInfra)
	swapUC := usecase.New(db)

	if err := server.Init(config.Server, swapUC); err != nil {
		logger.Log(common.Panic, err)
		panic(err)
	}
}
