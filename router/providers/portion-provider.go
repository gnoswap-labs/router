package providers

import (
	"router/router"
	"router/router/core"
	"router/router/core/entities/fractions"
)

// interface는 I 접두사를 붙이는 것이 관행인가?
type IPortionProvider interface {
	GetPortionAmount(tokenOutAmount fractions.CurrencyAmount, tradeType core.TradeType, swapConfig router.SwapOptions) fractions.CurrencyAmount
}

type PortionProvider struct {
}
