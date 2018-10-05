package server

import (
	"net/http"
)

//Settings struct is a simple settings for the server
type Settings struct {
	SSL  bool
	Port string `env:"ESS_ATOMIC_SWAP_APP_PORT"`
}

//Init is a func for initiate a Settings struct
func Init(settings Settings) {
	addr := ":" + settings.Port
	http.ListenAndServe(addr, nil)
}
