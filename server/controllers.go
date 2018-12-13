package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/essentiaone/ess-atomic-swap/models"
	"github.com/essentiaone/ess-atomic-swap/swap"
	"github.com/julienschmidt/httprouter"
)

type controllers struct {
	swap swap.AtomicSwapUseCase
}

func (c *controllers) intex(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	if _, err := w.Write([]byte("Hello, essentia!")); err != nil {
		log.Fatal("cannot write to index")
	}
}

func (c *controllers) initiate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	initiateInfo := &models.AtomicSwapInitiate{}
	err := json.NewDecoder(r.Body).Decode(&initiateInfo)
	if err != nil {
		generateResponse(w, "cannot create body response", false)
	}
	defer r.Body.Close()

	txHash, err := c.swap.Initiate(initiateInfo)
	if err != nil {
		generateResponse(w, err, false)
		return
	}
	generateResponse(w, txHash, true)
}

func generateResponse(w http.ResponseWriter, message interface{}, isSuccess bool) {
	w.Header().Add("Content-type", "application/json")
	key := "result"
	if !isSuccess {
		key = "error"
	}
	res := map[string]interface{}{
		key: message,
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Fatal("cannot write initial response")
	}
}
