package math

import (
	"errors"
	"math/big"
)

type Fraction struct {
	Numerator   *big.Int
	Denominator *big.Int
}

func NewFraction(numerator int64, denominator int64) *Fraction {
	if denominator == 0 {
		//return nil, errors.New("분모는 0일 수 없습니다")
		panic(errors.New("denominator is zero"))
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

func tryParseFraction(value interface{}) (*Fraction, error) {
	switch v := value.(type) {
	case *Fraction:
		return v, nil
	case *big.Int:
		return &Fraction{Numerator: v, Denominator: big.NewInt(1)}, nil
	case int64:
		return NewFraction(v, 1), nil
	default:
		return nil, errors.New("not a fraction")
	}
}

// 덧셈
func (f *Fraction) Add(other interface{}) *Fraction {
	o, err := tryParseFraction(other)
	if err != nil {
		panic("aaaaaaa")
	}

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

// 뺄셈
// TODO: 덧셈 활용해서 코드 단순화 할 수도 있을지도...
func (f *Fraction) Sub(other interface{}) *Fraction {
	o, err := tryParseFraction(other)
	if err != nil {
		panic("aaaaaa")
	}

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

func (f *Fraction) Mul(other *Fraction) *Fraction {
	g := GCD(f.Denominator, other.Denominator)

	numerator := new(big.Int).Mul(f.Numerator, other.Numerator)
	denominator := new(big.Int).Mul(f.Denominator, other.Denominator)

	return &Fraction{
		Numerator:   new(big.Int).Div(numerator, g),
		Denominator: new(big.Int).Div(denominator, g),
	}
}

func (f *Fraction) Div(other *Fraction) *Fraction {
	g := GCD(f.Denominator, other.Denominator)

	numerator := new(big.Int).Mul(f.Numerator, other.Denominator)
	denominator := new(big.Int).Mul(f.Denominator, other.Numerator)

	return &Fraction{
		Numerator:   new(big.Int).Div(numerator, g),
		Denominator: new(big.Int).Div(denominator, g),
	}
}

func (f *Fraction) LessThan(other *Fraction) bool {
	leftValue := new(big.Int).Mul(f.Numerator, other.Denominator)
	rightValue := new(big.Int).Mul(other.Numerator, f.Denominator)
	return leftValue.Cmp(rightValue) < 0
}

func (f *Fraction) GreaterThan(other *Fraction) bool {
	leftValue := new(big.Int).Mul(f.Numerator, other.Denominator)
	rightValue := new(big.Int).Mul(other.Numerator, f.Denominator)
	return leftValue.Cmp(rightValue) > 0
}

// 약분이 완료되었다는 가정
func (f *Fraction) Equals(other *Fraction) bool {
	return f.Numerator.Cmp(other.Numerator) == 0 &&
		f.Denominator.Cmp(other.Denominator) == 0
}

// 역
func (f *Fraction) Invert() *Fraction {
	return &Fraction{
		Numerator:   f.Denominator,
		Denominator: f.Numerator,
	}
}

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
