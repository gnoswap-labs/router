package currency

// Currency는 Token | NativeCurrency
type ICurrency interface {
	Equals(other ICurrency) bool
	GetToken() Token
	IsToken() bool
}

type Currency struct {
	IsNative bool
	IsToken  bool
	ChainId  int
	Decimals int
	Symbol   string
	Name     string
	address  string
}

func NewCurrency(chainId int, decimals int, symbol string, name string) *Currency {
	//invariant(Number.isSafeInteger(chainId), 'CHAIN_ID')
	//invariant(
	//	Decimals >= 0 && Decimals < 255 && Number.isInteger(Decimals),
	//	'DECIMALS',
	//)

	// 이 제약이 걸린 이유는 아직 알지 못한다.
	if decimals < 0 || 255 < decimals {
		panic("Decimals must be between 0 and 255")
	}

	return &Currency{
		ChainId:  chainId,
		Decimals: decimals,
		Symbol:   symbol,
		Name:     name,
	}
}
