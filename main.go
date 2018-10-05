package main

import (
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received Request: ", r.URL.Path)
	w.Write([]byte("Hello, world!"))
}
func main() {
	http.HandleFunc("/", handler)

	port := os.Getenv("ESS_ATOMIC_SWAP_APP_PORT")
	if port == "" {
		log.Fatal("ESS_ATOMIC_SWAP_APP_PORT environment variable was not set")
	}

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Could not listen: ", err)
	}
}
