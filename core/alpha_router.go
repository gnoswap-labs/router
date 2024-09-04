package core

import "router/core/currency"

type AlphaRouter struct {
	//chainId ChainId
	portionProvider IPortionProvider
}

func NewAlphaRouter(params AlphaRouterParams) *AlphaRouter {
	return &AlphaRouter{}
}

// Todo: goroutine
func (a *AlphaRouter) route(
	baseCurrency currency.Currency,
	quoteCurrency currency.Currency,
	amount float64, // todo: float64 -> fraction
	tradeType TradeType,
	swapConfig SwapOptions,
	routerConfig AlphaRouterConfig,
) SwapRoute {
	//originalAmount := amount // for save

	//currencyIn, currencyOut := a.determineCurrencyInOutFromTradeType(tradeType, baseCurrency, quoteCurrency)

	// token은 currency의 wrapped된 버전이다.
	//tokenIn := currencyIn.GetToken()
	//tokenOut := currencyOut.GetToken()

	// 왠만하면 함수로 뺄 것
	// 내용 이해 필요
	if tradeType == EXACT_OUTPUT {
		portionAmount, portionErr := a.portionProvider.GetPortionAmount(amount, tradeType, swapConfig)

		if portionErr == nil && portionAmount > 0 {
			// In case of exact out swap, before we route, we need to make sure that the
			// token out amount accounts for flat portion, and token in amount after the best swap route contains the token in equivalent of portion.
			// In other words, in case a pool's LP fee bps is lower than the portion bps (0.01%/0.05% for v3), a pool can go insolvency.
			// This is because instead of the swapper being responsible for the portion,
			// the pool instead gets responsible for the portion.
			// The addition below avoids that situation.
			amount += portionAmount
		}
	}

	// routing config merge다루는 부분 패스
	//routerConfig = setRouterConfig(routingConfig, chainId)

	// tokenIn 또는 tokenOut과 동일한 값...
	//quoteToken := quoteCurrency.GetToken()

	// main logic?
	//routes := a.getSwapRouteFromChain(tokenIn, tokenOut, amount, tradeType, routingConfig)

	//if routes == nil {
	//	// todo: error 처리 해 줄 것
	//}

	//trade := a.buildTrade(currencyIn, currencyOut, tradeType, routes)

	swapRoute := a.buildSwapRoute()
	return swapRoute
}

func (a *AlphaRouter) determineCurrencyInOutFromTradeType(
	tradeType TradeType,
	baseCurrency currency.Currency,
	quoteCurrency currency.Currency,
) (currency.Currency, currency.Currency) {
	if tradeType == EXACT_INPUT {
		return baseCurrency, quoteCurrency
	}
	return quoteCurrency, baseCurrency
}

// todo: goroutine
func (a *AlphaRouter) getSwapRouteFromChain(tokenIn, tokenOut currency.Token, amount float64, tradeType TradeType, routingConfig AlphaRouterConfig) *BestSwapRoute {
	//percents, amount := a.getAmountDistribution(amount, routingConfig)

	return &BestSwapRoute{}
}

func (a *AlphaRouter) getAmountDistribution(amount float64, routingConfig AlphaRouterConfig) (float64, float64) {

	return 0, 0
}

func (a *AlphaRouter) buildTrade(currencyIn currency.Currency, currencyOut currency.Currency, tradeType TradeType, routes Routes) Trade {

	return Trade{}
}

func (a *AlphaRouter) buildSwapRoute() SwapRoute {
	return SwapRoute{}
}

func (a *AlphaRouter) setRouterConfig(routerConfig AlphaRouterConfig, chainId int) AlphaRouterConfig {
	return AlphaRouterConfig{}
}
