package server

import (
	"github.com/essentiaone/ess-atomic-swap/swap"
	"github.com/julienschmidt/httprouter"
)

func initRouter(swap swap.AtomicSwapUseCase) *httprouter.Router {
	router := httprouter.New()

	controllers := &controllers{swap}

	router.GET("/", controllers.intex)
	router.POST("/initiate", controllers.initiate)

	return router
}
