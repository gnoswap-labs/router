package fractions

import (
	"errors"
	"math/big"
)

type Fraction struct {
	Numerator   *big.Int
	Denominator *big.Int
}

func NewFraction(numerator int64, denominator int64) *Fraction {
	return &Fraction{
		Numerator:   big.NewInt(numerator),
		Denominator: big.NewInt(denominator),
	}
}

func tryParseFraction(value interface{}) (Fraction, error) {
	switch v := value.(type) {
	case Fraction:
		return v, nil
	case *big.Int:
		return Fraction{Numerator: v, Denominator: big.NewInt(1)}, nil
	case big.Int:
		return Fraction{Numerator: &v, Denominator: big.NewInt(1)}, nil
	// TODO: int64 등 추가 바람
	default:
		return Fraction{}, errors.New("not a fraction")
	}
}

// ---------------------------
// TODO: 형을 자유롭게 받은 다음 tryParseFraction으로 한 번 거르고 사용하도록 하기
// 덧셈
func (f Fraction) Add(other *Fraction) *Fraction {
	numerator1 := new(big.Int).Mul(f.Numerator, other.Denominator)
	numerator2 := new(big.Int).Mul(other.Numerator, f.Denominator)
	denominator := new(big.Int).Mul(f.Denominator, other.Denominator)
	resultNumerator := new(big.Int).Add(numerator1, numerator2)

	// TODO: 약분
	//g := utils.Gcd(resultNumerator, denominator)

	return &Fraction{
		Numerator:   resultNumerator,
		Denominator: denominator,
	}
}

// 뺄셈
// TODO: 덧셈 활용해서 코드 단순화 할 수도 있을지도...
func (f Fraction) Sub(other *Fraction) *Fraction {
	numerator1 := new(big.Int).Mul(f.Numerator, other.Denominator)
	numerator2 := new(big.Int).Mul(other.Numerator, f.Denominator)
	denominator := new(big.Int).Mul(f.Denominator, other.Denominator)
	resultNumerator := new(big.Int).Sub(numerator1, numerator2)

	// TODO: 약분
	//g := utils.Gcd(resultNumerator, denominator)

	return &Fraction{
		Numerator:   resultNumerator,
		Denominator: denominator,
	}
}

func (f Fraction) LessThan(other *Fraction) bool {
	leftValue := new(big.Int).Mul(f.Numerator, other.Denominator)
	rightValue := new(big.Int).Mul(other.Numerator, f.Denominator)
	return leftValue.Cmp(rightValue) < 0
}

func (f Fraction) GreaterThan(other *Fraction) bool {
	leftValue := new(big.Int).Mul(f.Numerator, other.Denominator)
	rightValue := new(big.Int).Mul(other.Numerator, f.Denominator)
	return leftValue.Cmp(rightValue) > 0
}

// 약분이 완료되었다는 가정
// WARN: 약분이 안되었다는 가정이라면 약분하는 로직 추가해야 함
func (f Fraction) Equals(other *Fraction) bool {
	return f.Numerator.Cmp(other.Numerator) == 0 &&
		f.Denominator.Cmp(other.Denominator) == 0
}

// 몫
func (f Fraction) Quotient() *big.Int {
	return new(big.Int).Quo(f.Numerator, f.Denominator)
}

// NOTE: 단순 나머지가 아니라 분수 형태로 표시한다
// 좋은 형태의 함수는 아니라 생각
func (f Fraction) Remainder() Fraction {
	return Fraction{
		Numerator:   new(big.Int).Rem(f.Numerator, f.Denominator),
		Denominator: f.Denominator,
	}
}

// 역
func (f Fraction) Invert() Fraction {
	return Fraction{
		Numerator:   f.Denominator,
		Denominator: f.Numerator,
	}
}

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
