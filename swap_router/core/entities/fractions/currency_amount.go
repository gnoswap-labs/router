package fractions

import (
	"math/big"
	"router/swap_router/core"
	currency2 "router/swap_router/core/entities/currency"
	"router/swap_router/core/entities/fractions/math"
)

// TODO: 어떻게 구현할 것인가?
//type BigNumberish struct {
//}

type CurrencyAmount[T currency2.Currency] struct {
	math.Fraction
	currency     T
	decimalScale *big.Int
}

func NewCurrencyAmount[T currency2.Currency](currency T, numerator *big.Int, denominator *big.Int) *CurrencyAmount[T] {
	//invariant(JSBI.lessThanOrEqual(this.quotient, MaxUint256), 'AMOUNT');
	currencyAmount := CurrencyAmount[T]{
		Fraction: math.Fraction{
			Numerator:   numerator, // warn: pointer 받은 거 그대로 넣는게 맞나
			Denominator: denominator,
		},
		currency:     currency,
		decimalScale: big.NewInt(1),
	}

	return &currencyAmount
}

func (c *CurrencyAmount[T]) fromRawAmount(currency T, rawAmount *big.Int) *CurrencyAmount[T] {
	return NewCurrencyAmount(currency, rawAmount, big.NewInt(1))
}

func (c *CurrencyAmount[T]) fromFractionalAmount(currency T, numerator *big.Int, denominator *big.Int) *CurrencyAmount[T] {
	return NewCurrencyAmount(currency, numerator, denominator)
}

func (c *CurrencyAmount[T]) Add(other CurrencyAmount[T]) *CurrencyAmount[T] {
	added := c.Fraction.Add(&other.Fraction)
	return c.fromFractionalAmount(c.currency, added.Numerator, added.Denominator)
}

func (c *CurrencyAmount[T]) Sub(other CurrencyAmount[T]) *CurrencyAmount[T] {
	subbed := c.Fraction.Sub(&other.Fraction)
	return c.fromFractionalAmount(c.currency, subbed.Numerator, subbed.Denominator)
}

func (c *CurrencyAmount[T]) Multiply(other CurrencyAmount[T]) *CurrencyAmount[T] {
	multiplied := c.Fraction.Mul(&other.Fraction)
	return c.fromFractionalAmount(c.currency, multiplied.Numerator, multiplied.Denominator)
}

func (c *CurrencyAmount[T]) Divide(other CurrencyAmount[T]) *CurrencyAmount[T] {
	divided := c.Fraction.Div(&other.Fraction)
	return c.fromFractionalAmount(c.currency, divided.Numerator, divided.Denominator)
}

// 보류
func (c *CurrencyAmount[T]) ToSignificant(significantDigits int, rounding core.Rounding) string {
	return c.Fraction.Div(math.NewFraction(c.decimalScale, big.NewInt(1))).ToSignificant(significantDigits, rounding)
}

// 보류
func (c *CurrencyAmount[T]) ToFixed(decimalPlaces int, format interface{}, rounding core.Rounding) string {
	return ""
}

func (c *CurrencyAmount[T]) ToExact(format interface{}) string {
	return ""
}

func (c *CurrencyAmount[T]) GetToken() currency2.Token {
	if c.currency.IsToken {
		c.currency

	}

	return CurrencyAmount.fromFractionalAmount(c.currency.GetToken(), c.Numerator, c.Denominator)
}
