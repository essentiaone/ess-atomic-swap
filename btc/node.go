package btc

// BitcoinNode is an interface for speak with Bitcoin node
type BitcoinNode interface {
	ExecuteRequest(method string, target interface{}, params ...interface{}) (interface{}, error)
}
