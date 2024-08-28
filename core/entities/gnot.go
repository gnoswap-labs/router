package entities

import "router/core/currency"

type Gnot struct {
	NativeCurrency
}

func NewGnot(chainId int) *Gnot {
	return &Gnot{
		NativeCurrency: NativeCurrency{
			currency.BaseCurrency: currency.BaseCurrency{
				chainId: chainId,
			},
		},
	}
}

func (g Gnot) Wrapped() currency.Token {
	wgnot := WGNOT[g.chainId]
	// invariant(!!wgnot, 'WRAPPED')
	return wgnot
}
