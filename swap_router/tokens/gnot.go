package tokens

import (
	currency2 "router/swap_router/core/entities/currency"
)

type Gnot struct {
	currency2.NativeCurrency
}

func NewGnot(chainId1 int64) *Gnot {
	return &Gnot{
		NativeCurrency: currency2.NativeCurrency{
			Currency: currency2.Currency{
				ChainId: chainId1,
			},
		},
	}
}
