package server

import (
	"net/http"

	"github.com/essentiaone/ess-atomic-swap/swap"
)

// Settings struct is a simple settings for the server
type Settings struct {
	SSL  bool
	Port string `env:"PORT,required"`
}

// Init is a func for initiate a Settings struct
func Init(settings Settings, swap swap.AtomicSwapUseCase) error {
	addr := ":" + settings.Port
	routers := initRouter(swap)
	return http.ListenAndServe(addr, routers)
}
