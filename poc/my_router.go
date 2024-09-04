package poc

import (
	"fmt"
	"math"
)

// MyRouter
// router PoC
type MyRouter struct {
	network map[string]*Pool
	adj     map[string][]string
}

func NewMyRouter(edges []*Pool) *MyRouter {
	router := &MyRouter{
		network: make(map[string]*Pool),
		adj:     make(map[string][]string),
	}

	for _, edge := range edges {
		router.network[edge.Address] = edge
		router.adj[edge.TokenA.Symbol] = append(router.adj[edge.TokenA.Symbol], edge.TokenB.Symbol)
		router.adj[edge.TokenB.Symbol] = append(router.adj[edge.TokenB.Symbol], edge.TokenA.Symbol)
	}

	return router
}

// Route
// 두 개의 토큰 사이의 효율적인 경로를 계산하는 함수
func (m *MyRouter) Route(request SwapRequest) ([]SwapResult, error) {
	// V1 Router
	return m.findRouteV1(request)

	// V2 Router
	//return m.findRouteV2(startTokenSymbol, endTokenSymbol, AmountIn, 1)
}

// findRouteV1
// 두 토큰을 direct swap한다
func (m *MyRouter) findRouteV1(request SwapRequest) ([]SwapResult, error) {
	return m.swap(request.FromTokenSymbol, request.ToTokenSymbol, request.AmountIn)
}

// findRouteV2
// 경로가 maxLength 이하의 길이인 경로를 탐색해 route를 구한다
func (m *MyRouter) findRouteV2(request SwapRequest, maxLength int, routes []SwapResult) ([]SwapResult, error) {
	startTokenSymbol, beforeTokenSymbol, amountIn := m.setSymbolAndAmountIn(request, routes)
	if startTokenSymbol == request.ToTokenSymbol {
		return routes, nil
	}
	if len(routes) >= maxLength {
		return nil, fmt.Errorf("the length of routes exceeds maxLength")
	}

	var bestPath []SwapResult
	for _, toTokenSymbol := range m.adj[startTokenSymbol] {
		if toTokenSymbol == beforeTokenSymbol { // 경로 2인 cycle은 허용하지 않음
			continue
		}

		route, swapErr := m.swap(startTokenSymbol, toTokenSymbol, amountIn)
		if swapErr != nil {
			continue
		}

		workablePath, findRouteErr := m.findRouteV2(request, maxLength, append(routes, route...))
		if findRouteErr != nil {
			continue
		}

		if len(workablePath) != 0 && (bestPath == nil || (bestPath[len(bestPath)-1].AmountOut < workablePath[len(workablePath)-1].AmountOut)) {
			bestPath = workablePath
		}
	}

	return bestPath, nil
}

// setSymbolAndAmountIn
func (m *MyRouter) setSymbolAndAmountIn(request SwapRequest, routes []SwapResult) (string, string, float64) {
	if routes == nil { // 처음 함수가 호출된 거라면
		return request.FromTokenSymbol, "", request.AmountIn
	}
	return routes[len(routes)-1].OutTokenSymbol, routes[len(routes)-1].InTokenSymbol, routes[len(routes)-1].AmountOut
}

// Swap
// 두 개의 토큰 사이의 직접적인 Pool을 통해 두 개의 토큰을 교환하는 함수
func (m *MyRouter) swap(fromTokenSymbol string, toTokenSymbol string, amountIn float64) ([]SwapResult, error) {
	// TODO: poolName은 from:to가 아니라 to:from일 수 있다.
	// TODO: 문자열 연산 최적화
	poolName := fromTokenSymbol + ":" + toTokenSymbol

	if pool, ok := m.network[poolName]; ok {
		reserveFromToken, reserveToToken := m.getReserveOfTokenFromPool(fromTokenSymbol, toTokenSymbol, *pool)
		amountOut := m.calculateAmountOfToToken(reserveFromToken, reserveToToken, amountIn, *pool)

		// 같은 경로를 두 번 이상 탐색하지 않으므로 일단 주석 처리
		//m.saveSwap(fromTokenSymbol, amountIn, amountOut, pool)

		return []SwapResult{{
			InTokenSymbol:  fromTokenSymbol,
			OutTokenSymbol: toTokenSymbol,
			AmountIn:       amountIn,
			AmountOut:      math.Abs(amountOut),
		}}, nil
	}

	return nil, fmt.Errorf("pool %s not found", poolName)
}

// saveSwap
func (m *MyRouter) saveSwap(fromTokenSymbol string, amountIn, amountOut float64, pool *Pool) {
	if pool.TokenA.Symbol == fromTokenSymbol {
		pool.ReserveA += amountIn
		pool.ReserveB += amountOut
	} else {
		pool.ReserveA += amountOut
		pool.ReserveB += amountIn
	}
}

// getReserveOfTokenFromPool
// Pool에 있는 fromToken과 toToken의 reserve 쌍을 반환하는 함수
func (m *MyRouter) getReserveOfTokenFromPool(fromTokenSymbol string, toTokenSymbol string, pool Pool) (float64, float64) {
	if fromTokenSymbol == pool.TokenA.Symbol {
		return pool.ReserveA, pool.ReserveB
	}
	return pool.ReserveB, pool.ReserveA
}

// calculateAmountOfToToken
// 토큰이 교환될 때 교환자에게 지급해야 할 toToken의 양을 계산하는 함수
// 계산 과정 최적화 하면 곱셈 5번, 덧셈 2번 정도에 해결 가능함
// ref: https://hyun-jeong.medium.com/uniswap-series-2-cpmm-%EC%9D%B4%ED%95%B4%ED%95%98%EA%B8%B0-4a82de8aba9
func (m *MyRouter) calculateAmountOfToToken(reserveFromToken, reserveToToken, amountIn float64, pool Pool) float64 {
	X := reserveFromToken
	Y := reserveToToken
	dX := amountIn

	K := X * Y
	L := math.Sqrt(K)
	P := X / Y

	X_ := X + dX
	P_ := (X_ / L) * (X_ / L)

	dY := L * (1/math.Sqrt(P_) - 1/math.Sqrt(P))

	// X 코인이 dX개 만큼 증가했을 때
	// Y 코인은 dY개 만큼 감소해야 한다.
	// X -> X + dX, Y -> Y + dY
	return dY
}
