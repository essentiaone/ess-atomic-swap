package server

import (
	"net/http"
)

//Settings struct is a simple settings for the server
type Settings struct {
	SSL  bool
	Port string `env:"PORT"`
}

func initRouting() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("Hello, world")); err != nil {
			panic(err)
		}
	})
}

//Init is a func for initiate a Settings struct
func Init(settings Settings) error {
	addr := ":" + settings.Port
	initRouting()
	return http.ListenAndServe(addr, nil)
}
