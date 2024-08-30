package core

import "router/core/currency"

type AlphaRouter struct {
	//chainId ChainId
	//portionProvider IPortionProvider
}

func NewAlphaRouter(params AlphaRouterParams) *AlphaRouter {
	return &AlphaRouter{}
}

func (a AlphaRouter) route(
	baseCurrency currency.Currency, // currencyIn
	quoteCurrency currency.Currency, // currencyOut으로 바꿔도 될 것 같다.
	amount float64,
	// amount fractions.CurrencyAmount,
	// tradeType TradeType,
	// swapConfig SwapOptions,
) SwapRoute {
	//originalAmount := amount

	// currencyIn, currencyOut은 Currency 타입이고
	// Currency 타입은 NativeCurrency(GNOT)이거나 Token 타입이다.
	// 아래에서 Token 타입이길 원하는 듯하다.
	//tokenIn := currencyIn.Wrapped()
	//tokenOut := currencyOut.Wrapped()

	//core 패키지를 TradeType 패키지로 변경하면 가독성이 더 좋아질 듯 하다.
	//if tradeType == EXACT_OUTPUT {
	//	// TODO: GetPortionAmount에서 반환 값인 CurrencyAmount을 반환하지 못할 경우가 있을 수도 있다.(높은 확률로)
	//	portionAmount := a.portionProvider.GetPortionAmount(
	//		amount,
	//		tradeType,
	//		swapConfig,
	//	)
	//
	//result := portionAmount.GreaterThan(0)
	//if result {
	//	amount = amount.add(portionAmount)
	//}
	//}

	return SwapRoute{}
}

//
//func (a AlphaRouter) determineCurrencyInOutFromTradeType(
//	tradeType TradeType,
//	amount fractions.CurrencyAmount,
//	quoteCurrency currency.Currency,
//) (currency.Currency, currency.Currency) {
//	if tradeType == EXACT_INPUT {
//		return amount.Currency, quoteCurrency
//	} else {
//		return quoteCurrency, amount.Currency
//	}
//}
