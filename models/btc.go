package models

// BitcoinTransaction is a struct for a regular tx
type BitcoinTransaction struct {
	Amount        uint32
	Confirmations uint32
	TxID          uint32
	Time          uint32
	Details       struct {
		Account  string
		Address  string
		Category string
		Amount   uint32
		Fee      uint32
	}
}
