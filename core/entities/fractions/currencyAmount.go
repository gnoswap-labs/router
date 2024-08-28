package fractions

import (
<<<<<<< Updated upstream:core/entities/fractions/currencyAmount.go
	"router/core/entities"
)

type CurrencyAmount struct {
	Fraction
	Currency entities.Currency
=======
	"router/core/currency"
	"router/core/math"
)

type CurrencyAmount struct {
	math.Fraction
	Currency currency.Currency
>>>>>>> Stashed changes:core/entities/fractions/currency_amount.go
}

//func (c CurrencyAmount) add(other CurrencyAmount) CurrencyAmount {
//	added := c.Fraction.add(other)
//
//	return CurrencyAmount.fromFractionAmount()
//}
