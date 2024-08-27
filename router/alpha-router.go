package router

import (
	"router/router/core"
	"router/router/core/entities"
	"router/router/core/entities/fractions"
)

type AlphaRouter struct {
}

func NewAlphaRouter(params AlphaRouterParams) *AlphaRouter {
	return &AlphaRouter{}
}

// TODO: 원본 코드에서는 async 함수
// 라우트 한 결과는 SwapRoute
func (a AlphaRouter) route(
	amount fractions.CurrencyAmount,
	quoteCurrency entities.Currency,
	tradeType core.TradeType,
) SwapRoute {
	originalAmount := amount

	currencyIn, currencyOut := a.determineCurrencyInOutFromTradeType(tradeType, amount, quoteCurrency)

	// currencyIn, currencyOut은 Currency 타입이고
	// Currency 타입은 NativeCurrency이거나 Token 타입이다.
	// 아래에서 Token 타입이길 원하는 듯하다.
	tokenIn := currencyIn
	tokenOut := currencyOut

	swapRoute := SwapRoute{}
	return swapRoute
}

func (a AlphaRouter) determineCurrencyInOutFromTradeType(
	tradeType core.TradeType,
	amount fractions.CurrencyAmount,
	quoteCurrency entities.Currency,
) (entities.Currency, entities.Currency) {
	if tradeType == core.EXACT_INPUT {
		return amount.Currency, quoteCurrency
	} else {
		return quoteCurrency, amount.Currency
	}
}
