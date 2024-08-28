package core

// interface는 I 접두사를 붙이는 것이 관행인가?
type IPortionProvider interface {
	// GetPortionAmount(tokenOutAmount fractions.CurrencyAmount, tradeType TradeType, swapConfig SwapOptions) fractions.CurrencyAmount
}

type PortionProvider struct {
}
