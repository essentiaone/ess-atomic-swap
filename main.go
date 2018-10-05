package main

import (
	"log"
	"net/http"

	"github.com/essentiaone/ess-atomic-swap/config"
	"github.com/essentiaone/ess-atomic-swap/server"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received Request: ", r.URL.Path)
	w.Write([]byte("Hello, worrrrld!"))
}
func main() {

	settings, _ := config.Init()
	server.Init(settings.Server)
	log.Fatal(http.ListenAndServe(":"+settings.Server.Port, nil))
}
