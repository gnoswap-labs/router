package swap_router

import (
	"router/swap_router/core/entities/currency"
	"router/swap_router/v3_sdk"
)

type Pool struct {
	tokenA           currency.Token
	tokenB           currency.Token
	fee              v3_sdk.FeeAmount
	sqrtRatioX96     float64
	liquidity        float64
	tickCurrent      float64
	tickDataProvider TickDataProvider

	path string
}

func NewPool(tokenA currency.Token, tokenB currency.Token, fee FeeAmount, sqrtRatioX96 float64, liquidity float64, tickCurrent float64, tickDataProvider TickDataProvider) *Pool {

	return &Pool{}
}
