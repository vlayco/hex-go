package api

import "github.com/vlayco/hex-go/internal/ports"

type Adapter struct {
	// This adapter will contain the core adapter,
	// since api layer will access the core of the app.
	arith ports.ArithmeticPort
}
