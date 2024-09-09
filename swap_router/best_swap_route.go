package swap_router

import (
	"router/swap_router/core/entities/fractions/math"
)

type BestSwapRoute struct {
	quote               math.Fraction
	routeWithValidQuote []V3RouteWithValidQuote
}

// todo: goroutine
// 모든 경로에 대한 모든 금액에 대한 모든 견적이 주어지면 최상의 조합을 찾으십시오.
func getBestSwapRoute(amount math.Fraction, routesWithValidQuotes []V3RouteWithValidQuote, tradeType TradeType, routerConfig AlphaRouterConfig) *BestSwapRoute {
	// 유효한 인용문 목록에 대한 입력 비율 맵을 작성합니다.
	// 다양한 이유로(부족한 유동성 등) 견적이 null이 될 수 있으므로 여기에도 입력합니다.
	percentToQuotes := map[int][]V3RouteWithValidQuote{}

	// ????????
	for _, routeWithValidQuote := range routesWithValidQuotes {
		if !percentToQuotes[routeWithValidQuote.percent] {
			percentToQuotes[routeWithValidQuote.percent] = []V3RouteWithValidQuote{}
		}
		percentToQuotes[routeWithValidQuote.percent]!.push(routeWithValidQuote)
	}

	// 각 백분율에 대한 유효한 인용문이 모두 주어지면 최적의 경로를 찾습니다.
	bestSwapRoute := getBestSwapRouteBy()
	if bestSwapRoute == nil {
		// error
	}

	// 입력의 백분율을 계산할 때 정밀도가 손실될 수 있으므로 각 백분율의 합계가
	// 최적의 견적 경로가 정확한 In 또는 ExactOut에 정확히 합산되지 않을 수 있습니다.

	// 여기에서 확인하고 불일치가 있는 경우
	// 누락된 금액을 무작위 경로에 추가하세요. 누락된 금액 크기는 무시할 수 있으므로 견적은 여전히 매우 정확해야 합니다.


	return bestSwapRoute
}


// todo: goroutine
func getBestSwapRouteBy(tradeType TradeType) *BestSwapRoute {


	return &BestSwapRoute{}
}
