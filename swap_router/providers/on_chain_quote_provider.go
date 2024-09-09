package providers

import (
	"router/swap_router/core/entities/fractions"
)

type IOnChainQuoteProvider interface {
	GetQuotesManyExactIn(amountIns []fractions.CurrencyAmount, routes []TRoute, providerConfig ProviderConfig) OnChainQuotes
}
