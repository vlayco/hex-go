package api

import (
	"github.com/vlayco/hex-go/internal/ports"
)

// This is were our app access core logic.
// If we change something in our core, or if we entirely swap the
// implementation, as far as that code provides all of the methods
// defined in arithmetics' core interface, we're all good!

type Adapter struct {
	db ports.DBPort

	// This adapter will contain the core adapter,
	// since api layer will access the core of the app.
	// DI will be used in future at this point.
	arith ports.ArithmeticPort
}

func NewAdapter(db ports.DBPort, arith ports.ArithmeticPort) *Adapter {
	return &Adapter{db: db, arith: arith}
}

func (apiAdap Adapter) GetAddition(a, b int32) (int32, error) {
	res, err := apiAdap.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	err = apiAdap.db.AddToHistory(res, "addition")
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (apiAdap Adapter) GetSubstraction(a, b int32) (int32, error) {
	res, err := apiAdap.arith.Substraction(a, b)
	if err != nil {
		return 0, err
	}

	err = apiAdap.db.AddToHistory(res, "substraction")
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (apiAdap Adapter) GetMultiplication(a, b int32) (int32, error) {
	res, err := apiAdap.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}

	err = apiAdap.db.AddToHistory(res, "multiplication")
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (apiAdap Adapter) GetDivision(a, b int32) (int32, error) {
	res, err := apiAdap.arith.Division(a, b)
	if err != nil {
		return 0, err
	}

	err = apiAdap.db.AddToHistory(res, "division")
	if err != nil {
		return 0, err
	}

	return res, nil
}
