package eth

// EthereumNode is an interface for speak with ETH node
type EthereumNode interface {
	ExecuteRequest(method string, target interface{}, params ...interface{}) (interface{}, error)
}
