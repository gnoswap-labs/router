package core

import "router/core/currency"

type Pool struct {
	tokenA           currency.Token
	tokenB           currency.Token
	fee              FeeAmount
	sqrtRatioX96     float64
	liquidity        float64
	tickCurrent      float64
	tickDataProvider TickDataProvider

	path string
}

func NewPool(tokenA currency.Token, tokenB currency.Token, fee FeeAmount, sqrtRatioX96 float64, liquidity float64, tickCurrent float64, tickDataProvider TickDataProvider) *Pool {

	return &Pool{}
}
