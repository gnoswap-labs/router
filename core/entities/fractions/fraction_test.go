package fractions

import (
	"testing"
)

func TestFraction_Add(t *testing.T) {
	fraction1 := NewFraction(1, 2)
	fraction2 := NewFraction(1, 3)
	result := fraction1.Add(fraction2)
	expected := NewFraction(5, 6)

	if result.Numerator.Cmp(expected.Numerator) != 0 || result.Denominator.Cmp(expected.Denominator) != 0 {
		t.Fatalf("Add: expected %v, got %v", expected, result)
	}
}

func TestFraction_Sub(t *testing.T) {
	fraction1 := NewFraction(1, 2)
	fraction2 := NewFraction(1, 3)
	result := fraction1.Sub(fraction2)
	expected := NewFraction(1, 6)

	if result.Numerator.Cmp(expected.Numerator) != 0 || result.Denominator.Cmp(expected.Denominator) != 0 {
		t.Fatalf("Sub: expected %v, got %v", expected, result)
	}
}

func TestFraction_LessThan(t *testing.T) {
	fraction1 := NewFraction(1, 2)
	fraction2 := NewFraction(1, 3)
	result := fraction1.LessThan(fraction2)
	expected := false

	if result != expected {
		t.Fatalf("LessThan: expected %v, got %v", expected, result)
	}
}

func TestFraction_GreaterThan(t *testing.T) {
	fraction1 := NewFraction(1, 2)
	fraction2 := NewFraction(1, 3)
	result := fraction1.GreaterThan(fraction2)
	expected := true

	if result != expected {
		t.Fatalf("GreaterThan: expected %v, got %v", expected, result)
	}
}

func TestFraction_Equals(t *testing.T) {
	fraction1 := NewFraction(1, 2)
	fraction2 := NewFraction(1, 3)
	result := fraction1.Equals(fraction2)
	expected := false

	if result != expected {
		t.Fatalf("Equals: expected %v, got %v", expected, result)
	}
}
