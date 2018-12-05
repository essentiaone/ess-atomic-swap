package server

import (
	"encoding/json"
	"net/http"

	"github.com/essentiaone/ess-atomic-swap/models"
	"github.com/essentiaone/ess-atomic-swap/swap"
	"github.com/julienschmidt/httprouter"
)

type controllers struct {
	swap swap.AtomicSwapUseCase
}

func (c *controllers) intex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Hello, essentia!"))
}

func (c *controllers) initiate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	initiateInfo := &models.AtomicSwapInitiate{}
	err := json.NewDecoder(r.Body).Decode(&initiateInfo)
	defer func() {
		if r.Body != nil {
			r.Body.Close()
		}
	}()
	if err != nil {
		generateResponse(w, "cannot create body response", false)
	}
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
	rawRes, err := json.Marshal(res)
	if err != nil {
		w.Write([]byte(`{"error":"cannot create response body"}`))
		return
	}
	w.Write(rawRes)
	return
}
