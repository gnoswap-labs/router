package tokens

import "router/core/currency"

type Gnot struct {
	currency.NativeCurrency
}

func NewGnot(chainId1 int64) *Gnot {
	return &Gnot{
		NativeCurrency: currency.NativeCurrency{
			BaseCurrency: currency.BaseCurrency{
				ChainId: chainId1,
			},
		},
	}
}
