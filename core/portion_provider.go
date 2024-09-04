package core

type IPortionProvider interface {
	GetPortionAmount(tokenOutAmount float64, tradeType TradeType, swapConfig SwapOptions) (float64, error)
}

type PortionProvider struct {
}
