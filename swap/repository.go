package swap

// AtomicSwapRepository is an interface for storing info
type AtomicSwapRepository interface {
	SaveOrderTemporary() ([]byte, error)
	MoveOrderInvariable() ([]byte, error)
}
