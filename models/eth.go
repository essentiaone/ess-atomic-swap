package models

// EthereumTransaction is a struct for a regular tx
type EthereumTransaction struct {
	Logs              []interface{} `json:"logs"`
	TransactionHash   string        `json:"transactionHash"`
	TransactionIndex  string        `json:"transactionIndex"`
	BlockNumber       string        `json:"blockNumber"`
	BlockHash         string        `json:"blockHash"`
	CumulativeGasUsed string        `json:"cumulativeGasUse"`
	GasUsed           string        `json:"gasUsed"`
	ContractAddress   string        `json:"contractAddress"`
	LogsBloom         string        `json:"logsBloom"`
	Status            string        `json:"status"`
}
