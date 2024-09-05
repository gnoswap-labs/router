package core

import (
	"router/core/currency"
	"router/core/math"
)

type AlphaRouter struct {
	chainId         ChainId
	portionProvider IPortionProvider
}

func NewAlphaRouter(params AlphaRouterParams) *AlphaRouter {
	return &AlphaRouter{}
}

// Todo: goroutine
func (a *AlphaRouter) route(
	baseCurrency currency.Currency,
	quoteCurrency currency.Currency, // 여기는 왜 quote를 쓰지
	amount math.Fraction,
	tradeType TradeType,
	swapConfig SwapOptions,
	routerConfig AlphaRouterConfig,
) *SwapRoute {
	originalAmount := amount // for save

	currencyIn, currencyOut := a.determineCurrencyInOutFromTradeType(tradeType, baseCurrency, quoteCurrency)

	// token은 currency의 wrapped된 버전이다.
	tokenIn := currencyIn.GetToken()
	tokenOut := currencyOut.GetToken()

	// 왠만하면 함수로 뺄 것
	// TODO: 이해하기
	//if tradeType == EXACT_OUTPUT {
	//	portionAmount, portionErr := a.portionProvider.GetPortionAmount(amount, tradeType, swapConfig)
	//	if portionErr == nil && portionAmount > 0 {
	//		// 정확한 교환의 경우 라우팅하기 전에 다음 사항을 확인해야 합니다.
	//		// 토큰 아웃 금액은 고정 부분을 차지하며, 최상의 스왑 경로 이후의 토큰 인 금액에는 해당 부분에 해당하는 토큰이 포함됩니다.
	//		// 즉, 풀의 LP 수수료 bps가 부분 bps(v3의 경우 0.01%/0.05%)보다 낮은 경우 풀은 파산할 수 있습니다.
	//		// 그 이유는 스왑퍼가 해당 부분을 담당하는 대신,
	//		// 대신 풀이 해당 부분을 담당합니다.
	//		// 아래 추가 내용은 이러한 상황을 방지합니다.
	//		amount.Add(portionAmount)
	//	}
	//}

	// routing config merge다루는 부분 패스
	// 이름 리팩토링
	routerConfig = a.setRouterConfig(routerConfig)

	// tokenIn 또는 tokenOut과 동일한 값...
	//quoteToken := quoteCurrency.GetToken()

	// main logic?
	// todo: 원본 코드에서 quote, routes를 활용하는데 quote가 무엇인가?

	// 후처리는 나중으로
	bestSwapRoute := a.getSwapRouteFromChain(tokenIn, tokenOut, amount, tradeType, routerConfig)
	if bestSwapRoute == nil {
		return nil
	}

	trade := a.buildTrade(currencyIn, currencyOut, tradeType, bestSwapRoute.routeWithValidQuote)

	swapRoute := a.buildSwapRoute(trade)
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
// quoteToken Token을 getSwapRouteFromChain에 매개변수로 넘기면 계산 효율은 좋아진다(아주 미세하게), 하지만 quoteToken은 지금의 매개변수로 간단하게 알 수 있다.
func (a *AlphaRouter) getSwapRouteFromChain(tokenIn, tokenOut currency.Token, amount math.Fraction, tradeType TradeType, routerConfig AlphaRouterConfig) *BestSwapRoute {
	// 금액 분포, 즉 입력 금액의 일부를 생성합니다.
	// 다양한 경로에 대한 입력 금액의 일부에 대한 견적을 받은 다음
	// 결합하여 분할 경로를 생성합니다.
	// amount를 어떤 값들로(어떤 퍼센트 만큼씩) 나눌 것인지
	quotes := a.getAmountDistribution(amount, routerConfig.distributionPercent)

	// black box
	v3CandidatePools := getV3CandidatePools()
	// const quotePromises: Promisse<GetQuotesResult>[] = [];
	// const getQuotesResults = await Promise.all(quotePromises)

	// main logic
	// 모든 경로에 대한 모든 금액에 대한 모든 견적이 주어지면 최상의 조합을 찾으십시오.
	bestSwapRoute := getBestSwapRoute(amount, tradeType, routerConfig)

	return bestSwapRoute
}

func (a *AlphaRouter) getAmountDistribution(amount math.Fraction, distributionPercent int) []Quote {
	var quotes []Quote

	for i := 1; i <= 100; i += distributionPercent {
		quotes = append(quotes, Quote{
			percent: math.NewFraction(int64(i), 100),
			amount:  amount.Mul(math.NewFraction(int64(i), 100)),
		})
	}

	return quotes
}

func (a *AlphaRouter) buildTrade(currencyIn currency.Currency, currencyOut currency.Currency, tradeType TradeType, routes []V3RouteWithValidQuote) Trade {

	return Trade{}
}

func (a *AlphaRouter) buildSwapRoute(trade Trade) *SwapRoute {
	return &SwapRoute{}
}

// todo: 함수명이 명확하지 않다
func (a *AlphaRouter) setRouterConfig(routerConfig AlphaRouterConfig) AlphaRouterConfig {
	defaultConfig := AlphaRouterConfig{
		v3ProtocolPoolSelection: ProtocolPoolSelection{
			topN:                  2,
			topNDirectSwaps:       2,
			topNTokenInOut:        3,
			topNSecondHop:         1,
			topNWithEachBaseToken: 3,
			topNWithBaseToken:     5,
		},
		maxSwapsPerPath:        3,
		minSplits:              1,
		maxSplits:              7,
		distributionPercent:    5,
		useCachedRoutes:        true,
		writeToCachedRoutes:    true,
		optimisticCachedRoutes: false,
	}

	if routerConfig.v3ProtocolPoolSelection != (ProtocolPoolSelection{}) {
		defaultConfig.v3ProtocolPoolSelection = routerConfig.v3ProtocolPoolSelection
	}
	if routerConfig.maxSwapsPerPath != 0 {
		defaultConfig.maxSwapsPerPath = routerConfig.maxSwapsPerPath
	}
	if routerConfig.maxSplits != 0 {
		defaultConfig.maxSplits = routerConfig.maxSplits
	}
	if routerConfig.minSplits != 0 {
		defaultConfig.minSplits = routerConfig.minSplits
	}
	if routerConfig.distributionPercent != 0 {
		defaultConfig.distributionPercent = routerConfig.distributionPercent
	}
	return defaultConfig
}
