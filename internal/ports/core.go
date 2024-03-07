package ports

// internal/ports/core.go file houses the buisiness logic of our application.
// In this example app we are going to implement simple arithmetic operation,
// to keep it as simple as posible.

type ArithmeticPort interface {
	Addition(a int32, b int32) (int32, error)
	Substraction(a int32, b int32) (int32, error)
	Multiplication(a int32, b int32) (int32, error)
	Division(a int32, b int32) (int32, error)
}
