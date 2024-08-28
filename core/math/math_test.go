package math

import (
	"math/big"
	"testing"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		a, b, expected int64
	}{
		{1, 1, 1},
		{48, 18, 6},
		{101, 10, 1},
		{0, 5, 5},
		{5, 0, 5},
		{-48, 18, 6},
		{48, -18, 6},
		{-48, -18, 6},
		{16, -2, 2},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := GCD(big.NewInt(test.a), big.NewInt(test.b))

			if result.Cmp(big.NewInt(test.expected)) != 0 {
				t.Fatalf("GCD: expected %v, got %v", test.expected, result)
			}
		})
	}
}
