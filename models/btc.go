package models

// BitcoinDetailTransaction is a struct for a detail tx
type BitcoinDetailTransaction struct {
	Account  string  `json:"account"`
	Address  string  `json:"address"`
	Category string  `json:"category"`
	Amount   float32 `json:"amount"`
	Fee      float32 `json:"fee"`
}

// BitcoinTransaction is a struct for a regular tx
type BitcoinTransaction struct {
	Amount        float32                     `json:"amount"`
	Confirmations uint32                      `json:"confirmations"`
	TxID          string                      `json:"txid"`
	Time          uint32                      `json:"time"`
	Details       []*BitcoinDetailTransaction `json:"details"`
}
