package entities

type BaseCurrency struct {
	isNative bool
	isToken  bool

	chainId  int
	decimals int

	symbol  *string
	name    *string
	address *string
}

func NewBaseCurrency(chainId int, decimals int, symbol string, name string) *BaseCurrency {
	return &BaseCurrency{
		// 아래 코드는 원문
		//invariant(Number.isSafeInteger(chainId), 'CHAIN_ID');
		//invariant(
		//	decimals >= 0 && decimals < 255 && Number.isInteger(decimals),
		//	'DECIMALS',
		//);

		chainId:  chainId,
		decimals: decimals,
		symbol:   &symbol,
		name:     &name,
	}
}
