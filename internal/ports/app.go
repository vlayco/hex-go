package ports

// Ports are interfaces basicaly?
// And adapters provide implementations of those methods.

type APIPort interface {
	GetAddition(a, b int32) (int32, error)
	GetSubstraction(a, b int32) (int32, error)
	GetMultiplication(a, b int32) (int32, error)
	GetDivision(a, b int32) (int32, error)
}
