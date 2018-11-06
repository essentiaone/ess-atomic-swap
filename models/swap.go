package models

// AtomicSwapInitiate info for create atomic-swap order
type AtomicSwapInitiate struct {
	From     string  `json:"from"`
	To       string  `json:"to"`
	Password string  `json:"password"`
	Amount   float64 `josn:"amount"`
}
