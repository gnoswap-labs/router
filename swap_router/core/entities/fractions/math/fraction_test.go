package math

import (
	"math/big"
	"testing"
)

func TestNewFraction(t *testing.T) {
	tests := []struct {
		numerator, denominator                 int64
		expectedNumerator, expectedDenominator int64
		expectedPanic                          bool
	}{
		{1, 1, 1, 1, false},
		{2, 2, 1, 1, false},
		{1, 3, 1, 3, false},
		{1, 0, 0, 0, true},
		{6, 3, 2, 1, false},
		{12, 3, 4, 1, false},
		{-16, 2, -8, 1, false},
		{16, -2, 8, -1, false},
		{16, -2, -8, 1, false},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !test.expectedPanic {
						t.Fatalf("NewFraction: unexpected panic: %v", r)
					}
				}
			}()
			fraction := NewFraction(big.NewInt(test.numerator), big.NewInt(test.denominator))
			expectedFraction := NewFraction(big.NewInt(test.expectedNumerator), big.NewInt(test.expectedDenominator))

			if expectedFraction.Numerator.Cmp(fraction.Numerator) != 0 || expectedFraction.Denominator.Cmp(fraction.Denominator) != 0 {
				t.Fatalf("NewFraction: expected %v, got %v", expectedFraction, fraction)
			}
		})
	}
}

func TestFraction_Add(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                                   string
		numerator1, denominator1               int64
		numerator2, denominator2               int64
		expectedNumerator, expectedDenominator int64
	}{
		{"1/2 + 1/2", 1, 2, 1, 2, 1, 1},
		{"1/2 + 1/3", 1, 2, 1, 3, 5, 6},
		{"0/1 + 1/1", 0, 1, 1, 1, 1, 1},
		{"-1/2 + 1/2", -1, 2, 1, 2, 0, 1},
		{"1/-2 + 1/-3", 1, -2, 1, -3, -5, 6},
		{"2/1 + 15/7", 2, 1, 15, 7, 29, 7},
		{"2/1 + -15/7", 2, 1, -15, 7, -1, 7},
		{"1000000/1000001 + 1000000/1000001", 1000000, 1000001, 1000000, 1000001, 2000000, 1000001},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			fraction1 := NewFraction(big.NewInt(tt.numerator1), big.NewInt(tt.denominator1))
			fraction2 := NewFraction(big.NewInt(tt.numerator2), big.NewInt(tt.denominator2))
			result := fraction1.Add(fraction2)
			expected := NewFraction(big.NewInt(tt.expectedNumerator), big.NewInt(tt.expectedDenominator))

			if result.Numerator.Cmp(expected.Numerator) != 0 || result.Denominator.Cmp(expected.Denominator) != 0 {
				t.Fatalf("Add: expected %v, got %v", expected, result)
			}
		})
	}
}

func TestFraction_Sub(t *testing.T) {
	t.Parallel()

	tests := []struct {
		numerator1, denominator1               int64
		numerator2, denominator2               int64
		expectedNumerator, expectedDenominator int64
	}{
		{1, 2, 1, 3, 1, 6},
		{-1, 2, 1, 3, -5, 6},
		{0, 1, 16, 2, 8, -1},
		{1, 2, 1, 2, 0, 1},
		{1000000, 1000001, 1, 1000001, 999999, 1000001},
	}

	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			fraction1 := NewFraction(big.NewInt(tt.numerator1), big.NewInt(tt.denominator1))
			fraction2 := NewFraction(big.NewInt(tt.numerator2), big.NewInt(tt.denominator2))
			result := fraction1.Sub(fraction2)
			expected := NewFraction(big.NewInt(tt.expectedNumerator), big.NewInt(tt.expectedDenominator))

			if result.Numerator.Cmp(expected.Numerator) != 0 || result.Denominator.Cmp(expected.Denominator) != 0 {
				t.Fatalf("Sub: expected %v, got %v", expected, result)
			}
		})
	}
}

func TestFraction_Mul(t *testing.T) {
	t.Parallel()
	tests := []struct {
		numerator1, denominator1               int64
		numerator2, denominator2               int64
		expectedNumerator, expectedDenominator int64
		expectedPanic                          bool
	}{
		{1, 2, 1, 3, 1, 6, false},
		{-100, 10, 256, -10, 256, 1, false},
		{0, 1, 1, 1, 0, 1, false},
		{1, 2, 0, 1, 0, 1, false},
		{1, 2, 1, 0, 0, 0, true},
		{1, 2, -1, 2, -1, 4, false},
		{-1, 2, 1, 2, -1, 4, false},
		{1, 3, 1, 2, 1, 6, false},
		{2, 3, 3, 2, 1, 1, false},
		{2, 3, -3, 2, -1, 1, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tt.expectedPanic {
						return
					}
					t.Fatalf("Mul: unexpected panic: %v", r)
				}
			}()
			fraction1 := NewFraction(big.NewInt(tt.numerator1), big.NewInt(tt.denominator1))
			fraction2 := NewFraction(big.NewInt(tt.numerator2), big.NewInt(tt.denominator2))
			result := fraction1.Mul(fraction2)
			expected := NewFraction(big.NewInt(tt.expectedNumerator), big.NewInt(tt.expectedDenominator))

			if result.Numerator.Cmp(expected.Numerator) != 0 || result.Denominator.Cmp(expected.Denominator) != 0 {
				t.Fatalf("Mul: expected %v, got %v", expected, result)
			}
		})
	}
}

func TestFraction_Div(t *testing.T) {
	t.Parallel()
	tests := []struct {
		numerator1, denominator1               int64
		numerator2, denominator2               int64
		expectedNumerator, expectedDenominator int64
		expectedPanic                          bool
	}{
		{1, 2, 1, 3, 3, 2, false},
		{-100, 10, 256, -10, 100, 256, false},
		{0, 1, 1, 1, 0, 1, false},
		{1, 2, 1, 0, 0, 0, true},
		{-1, 2, 1, 2, -1, 1, false},
		{1, 3, 1, 2, 2, 3, false},
		{2, 3, 3, 2, 4, 9, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.expectedPanic {
						t.Fatalf("Div: unexpected panic: %v", r)
					}
				}
			}()
			fraction1 := NewFraction(big.NewInt(tt.numerator1), big.NewInt(tt.denominator1))
			fraction2 := NewFraction(big.NewInt(tt.numerator2), big.NewInt(tt.denominator2))
			result := fraction1.Div(fraction2)
			expected := NewFraction(big.NewInt(tt.expectedNumerator), big.NewInt(tt.expectedDenominator))

			if result.Numerator.Cmp(expected.Numerator) != 0 || result.Denominator.Cmp(expected.Denominator) != 0 {
				t.Fatalf("Div: expected %v/%v, got %v/%v", expected.Numerator, expected.Denominator, result.Numerator, result.Denominator)
			}
		})
	}
}

//func TestFraction_LessThan(t *testing.T) {
//	fraction1 := NewFraction(1, 2)
//	fraction2 := NewFraction(1, 3)
//	result := fraction1.LessThan(fraction2)
//	expected := false
//
//	if result != expected {
//		t.Fatalf("LessThan: expected %v, got %v", expected, result)
//	}
//}
//
//func TestFraction_GreaterThan(t *testing.T) {
//	fraction1 := NewFraction(1, 2)
//	fraction2 := NewFraction(1, 3)
//	result := fraction1.GreaterThan(fraction2)
//	expected := true
//
//	if result != expected {
//		t.Fatalf("GreaterThan: expected %v, got %v", expected, result)
//	}
//}
//
//func TestFraction_Equals(t *testing.T) {
//	fraction1 := NewFraction(1, 2)
//	fraction2 := NewFraction(1, 3)
//	result := fraction1.Equals(fraction2)
//	expected := false
//
//	if result != expected {
//		t.Fatalf("Equals: expected %v, got %v", expected, result)
//	}
//}
//

//
//func TestFraction_Invert(t *testing.T) {
//	fraction1 := NewFraction(1, 2)
//	result := fraction1.Invert()
//	expected := NewFraction(2, 1)
//
//	if result.Numerator.Cmp(expected.Numerator) != 0 || result.Denominator.Cmp(expected.Denominator) != 0 {
//		t.Fatalf("Invert: expected %v, got %v", expected, result)
//	}
//}
