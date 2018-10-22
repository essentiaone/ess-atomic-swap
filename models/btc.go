package models

//BTCTransaction is a struct for a regular tx
type BTCTransaction struct {
	Amount        uint32
	Confirmations uint32
	TxID          uint32
	Time          uint32
	Details       struct {
		account  string
		address  string
		category string
		amount   uint32
		fee      uint32
	}
}
