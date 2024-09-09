package swap_router

import (
	"router/swap_router/core/entities/fractions/math"
)

type IPortionProvider interface {
	GetPortionAmount(tokenOutAmount math.Fraction, tradeType TradeType, swapConfig SwapOptions) (math.Fraction, error)
}

type PortionProvider struct {
}
