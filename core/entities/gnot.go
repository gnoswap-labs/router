package entities

type Gnot struct {
	NativeCurrency
}

func NewGnot(chainId int) *Gnot {
	return &Gnot{
		NativeCurrency: NativeCurrency{
			BaseCurrency: BaseCurrency{
				chainId: chainId,
			},
		},
	}
}

func (g Gnot) Wrapped() Token {
	wgnot := WGNOT[g.chainId]
	// invariant(!!wgnot, 'WRAPPED')
	return wgnot
}
