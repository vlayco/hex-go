package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddition(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Addition(1, 1)
	if err != nil {
		t.Fatalf("exptected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(2))
}

func TestSubtraction(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Substraction(2, 1)
	if err != nil {
		t.Fatalf("exptected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(1))
}

func TestMultiplication(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Multiplication(2, 3)
	if err != nil {
		t.Fatalf("exptected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(6))
}

func TestDivision(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Division(9, 3)
	if err != nil {
		t.Fatalf("exptected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(3))
}
