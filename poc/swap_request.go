package poc

type SwapRequest struct {
	FromTokenSymbol string
	ToTokenSymbol   string
	AmountIn        float64
	//MinAmountOut int
	//UserAddress  string // option
}
