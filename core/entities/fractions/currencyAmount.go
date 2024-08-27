package fractions

import (
	"router/core/entities"
)

type CurrencyAmount struct {
	Fraction
	Currency entities.Currency
}

//func (c CurrencyAmount) add(other CurrencyAmount) CurrencyAmount {
//	added := c.Fraction.add(other)
//
//	return CurrencyAmount.fromFractionAmount()
//}
