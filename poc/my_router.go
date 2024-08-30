package poc

import (
	"fmt"
	"math"
)

type MyRouter struct {
	network map[string]*Pool
}

func NewMyRouter(edges []*Pool) *MyRouter {
	router := &MyRouter{
		network: make(map[string]*Pool),
	}

	for _, edge := range edges {
		router.network[edge.Address] = edge
	}

	return router
}

func (m *MyRouter) Swap(request SwapRequest) (SwapResult, error) {
	// poolName은 from:to가 아니라 to:from일 수 있다.
	poolName := request.FromToken + ":" + request.ToToken

	if pool, ok := m.network[poolName]; ok {
		fmt.Printf("pool found: %v\n", pool)

		reserveFromToken, reserveToToken := m.getReserveOfTokenFromPool(request.FromToken, request.ToToken, *pool)
		exchangedAmount := m.calculateAmountOfToToken(reserveFromToken, reserveToToken, request.AmountIn, *pool)

		//saveSwap()
		// TODO: 지금은 간이로 코드 작성하고 나중에 함수로 빼든 리팩토링 할 것
		if pool.TokenA.Symbol == request.FromToken {
			pool.ReserveA += request.AmountIn
			pool.ReserveB += exchangedAmount
		} else {
			pool.ReserveA += exchangedAmount
			pool.ReserveB += request.AmountIn
		}

		return SwapResult{
			AmountIn:  request.AmountIn,
			AmountOut: math.Abs(exchangedAmount),
		}, nil
	}

	return SwapResult{}, fmt.Errorf("pool %s not found", poolName)
}

func (m *MyRouter) getReserveOfTokenFromPool(fromTokenName string, toTokenName string, pool Pool) (float64, float64) {
	if fromTokenName == pool.TokenA.Symbol {
		return pool.ReserveA, pool.ReserveB
	}
	return pool.ReserveB, pool.ReserveA
}

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

func (m *MyRouter) dijskrtra() {

}
