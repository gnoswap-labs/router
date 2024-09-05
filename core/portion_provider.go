package core

import "router/core/math"

type IPortionProvider interface {
	GetPortionAmount(tokenOutAmount math.Fraction, tradeType TradeType, swapConfig SwapOptions) (math.Fraction, error)
}

type PortionProvider struct {
}
