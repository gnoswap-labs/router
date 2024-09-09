package currency

type Gnot struct {
	NativeCurrency
	_gnotCache map[int]Gnot
}

func NewGnot(chainId int) *Gnot {
	newGnot := &Gnot{}
	newGnot.Currency.ChainId = chainId
	newGnot.Currency.Decimals = 6
	newGnot.Currency.Symbol = "GNOT"
	newGnot.Currency.Name = "Gnot"
	return newGnot
}

func (g *Gnot) Wrapped() Token {
	wgnot := WGNOT[g.ChainId]
	//invariant(!!wgnot, 'WRAPPED');
	return wgnot
}

func (g *Gnot) onChain(chainId int) Gnot {
	if gnot, exists := g._gnotCache[chainId]; exists {
		return gnot
	}

	newGnot := Gnot{}
	newGnot.Currency.ChainId = chainId
	g._gnotCache[chainId] = newGnot
	return newGnot
}

func (g *Gnot) equals(other Currency) bool {
	return other.IsNative && (other.ChainId == g.ChainId)
}
