package arithmetic

// Adapters implement methods defined at ports interfaces
// Ports define interfaces to modules, and adapters provide,
// concrete implementations.

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (arith Adapter) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}

func (arith Adapter) Substraction(a int32, b int32) (int32, error) {
	return a - b, nil
}

func (arith Adapter) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}

func (arith Adapter) Division(a int32, b int32) (int32, error) {
	return a / b, nil
}
