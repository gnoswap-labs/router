package currency

type BaseCurrency struct {
	// The chain ID on which this currency resides
	chainId int
	// The decimals used in representing currency amounts
	decimals int

	// 이 아래 필드는 옵션
	// The symbol of the currency, i.e. a short textual non-unique identifier
	symbol string
	// The name of the currency, i.e. a descriptive textual non-unique identifier
	name    string
	address string
}

func NewBaseCurrency(chainId int, decimals int, symbol string, name string) *BaseCurrency {
	// 아래 코드는 원문
	//invariant(Number.isSafeInteger(chainId), 'CHAIN_ID');
	//invariant(
	//	decimals >= 0 && decimals < 255 && Number.isInteger(decimals),
	//	'DECIMALS',
	//);

	// 이 제약이 걸린 이유는 아직 알지 못한다.
	if decimals < 0 || 255 < decimals {
		panic("decimals must be between 0 and 255")
	}

	return &BaseCurrency{
		chainId:  chainId,
		decimals: decimals,
		symbol:   symbol,
		name:     name,
	}
}
