package fractions

import (
	"router/core/entities"
	"router/core/math"
)

type CurrencyAmount struct {
	math.Fraction
	Currency entities.Currency
}

//func (c CurrencyAmount) add(other CurrencyAmount) CurrencyAmount {
//	added := c.Fraction.add(other)
//
//	return CurrencyAmount.fromFractionAmount()
//}
