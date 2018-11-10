package models

// AtomicSwapInitiate data for initiate atomic-swap order
type AtomicSwapInitiate struct {
	From     string  `json:"from"`
	To       string  `json:"to"`
	Password string  `json:"password"`
	Amount   float64 `josn:"amount"`
}
