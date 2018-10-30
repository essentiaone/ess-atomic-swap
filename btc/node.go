package btc

// BitcoinNode is an interface for speak with BTC node
type BitcoinNode interface {
	ExecuteRequest(method string, target interface{}, params ...interface{}) (interface{}, error)
}
