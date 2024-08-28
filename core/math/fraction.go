package math

import (
	"errors"
	"math/big"
)

// Fraction
// 큰 수의 분수 표현을 위해 만든 구조체.
// 모든 수는 계산이 끝날 때마다 약분되며 0은 0/1로 표현한다.
// 또한 음수는 분자에만 표현된다. (2/(-3)은 -2/3으로 표현한다.)
type Fraction struct {
	Numerator   *big.Int
	Denominator *big.Int
}

// NewFraction
// int64 변수 2개를 분자, 분모 순으로 입력받아 Fraction 객체를 생성해주는 함수.
// 분모로 0이 입력된다면 panic을 발생시킨다.
func NewFraction(numerator int64, denominator int64) *Fraction {
	if denominator == 0 {
		panic(errors.New("denominator cannot be zero"))
	}

	// 분모는 양수만
	if denominator < 0 {
		numerator = -numerator
		denominator = -denominator
	}

	fraction := &Fraction{
		Numerator:   big.NewInt(numerator),
		Denominator: big.NewInt(denominator),
	}

	g := GCD(fraction.Numerator, fraction.Denominator)
	fraction.Numerator = fraction.Numerator.Div(fraction.Numerator, g)
	fraction.Denominator = fraction.Denominator.Div(fraction.Denominator, g)

	return fraction
}

// Add
// 앞의 분수에서 뒤의 분수를 더한 값을 리턴하는 함수.
func (f *Fraction) Add(o *Fraction) *Fraction {
	numerator1 := new(big.Int).Mul(f.Numerator, o.Denominator)
	numerator2 := new(big.Int).Mul(o.Numerator, f.Denominator)
	denominator := new(big.Int).Mul(f.Denominator, o.Denominator)
	resultNumerator := new(big.Int).Add(numerator1, numerator2)

	g := GCD(resultNumerator, denominator)

	return &Fraction{
		Numerator:   new(big.Int).Div(resultNumerator, g),
		Denominator: new(big.Int).Div(denominator, g),
	}
}

// Sub
// 앞의 분수에서 뒤의 분수를 뺀 값을 리턴하는 함수.
func (f *Fraction) Sub(o *Fraction) *Fraction {
	numerator1 := new(big.Int).Mul(f.Numerator, o.Denominator)
	numerator2 := new(big.Int).Mul(o.Numerator, f.Denominator)
	denominator := new(big.Int).Mul(f.Denominator, o.Denominator)
	resultNumerator := new(big.Int).Sub(numerator1, numerator2)

	g := GCD(resultNumerator, denominator)

	return &Fraction{
		Numerator:   new(big.Int).Div(resultNumerator, g),
		Denominator: new(big.Int).Div(denominator, g),
	}
}

// Mul
// 앞의 분수에서 뒤의 분수를 곱한 값을 리턴하는 함수.
func (f *Fraction) Mul(other *Fraction) *Fraction {
	numerator := new(big.Int).Mul(f.Numerator, other.Numerator)
	denominator := new(big.Int).Mul(f.Denominator, other.Denominator)

	if numerator.Sign() == -1 && denominator.Sign() == -1 {
		numerator.Neg(numerator)
		denominator.Neg(denominator)
	}

	g := GCD(numerator, denominator)

	return &Fraction{
		Numerator:   new(big.Int).Div(numerator, g),
		Denominator: new(big.Int).Div(denominator, g),
	}
}

// Div
// 앞의 분수에서 뒤의 분수를 나눈 값을 리턴하는 함수.
func (f *Fraction) Div(other *Fraction) *Fraction {
	numerator := new(big.Int).Mul(f.Numerator, other.Denominator)
	denominator := new(big.Int).Mul(f.Denominator, other.Numerator)

	if numerator.Sign() == -1 && denominator.Sign() == -1 {
		numerator.Neg(numerator)
		denominator.Neg(denominator)
	}

	g := GCD(numerator, denominator)

	return &Fraction{
		Numerator:   new(big.Int).Div(numerator, g),
		Denominator: new(big.Int).Div(denominator, g),
	}
}

//func (f *Fraction) LessThan(other *Fraction) bool {
//	leftValue := new(big.Int).Mul(f.Numerator, other.Denominator)
//	rightValue := new(big.Int).Mul(other.Numerator, f.Denominator)
//	return leftValue.Cmp(rightValue) < 0
//}
//
//func (f *Fraction) GreaterThan(other *Fraction) bool {
//	leftValue := new(big.Int).Mul(f.Numerator, other.Denominator)
//	rightValue := new(big.Int).Mul(other.Numerator, f.Denominator)
//	return leftValue.Cmp(rightValue) > 0
//}
//
//// 약분이 완료되었다는 가정
//func (f *Fraction) Equals(other *Fraction) bool {
//	return f.Numerator.Cmp(other.Numerator) == 0 &&
//		f.Denominator.Cmp(other.Denominator) == 0
//}
//
//// 역
//func (f *Fraction) Invert() *Fraction {
//	return &Fraction{
//		Numerator:   f.Denominator,
//		Denominator: f.Numerator,
//	}
//}

// 몫
//func (f Fraction) Quotient() *big.Int {
//	return new(big.Int).Quo(f.Numerator, f.Denominator)
//}

// NOTE: 단순 나머지가 아니라 분수 형태로 표시한다
// 좋은 형태의 함수는 아니라 생각
//func (f Fraction) Remainder() Fraction {
//	return Fraction{
//		Numerator:   new(big.Int).Rem(f.Numerator, f.Denominator),
//		Denominator: f.Denominator,
//	}
//}

//func (f Fraction) GreaterThan(other interface{}) (bool, error) {
//	otherParsed, err := tryParseFraction(other)
//	if err != nil {
//		return false, err
//	}
//left := f.Numerator * otherParsed.Denominator
//right := otherParsed.Denominator * f.Numerator
//
//	return left > right, nil
//}

// 구현 안 한 함수들
// toSignificant() {}
// toFixed() {}
