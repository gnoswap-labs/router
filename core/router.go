package core

type SwapRoute struct {
	route Routes
}

type IRouter interface {
}

// TODO: 원문: type SwapOptions = SwapOptionsUniversalRouter | SwapOptionsSwapRouter02
type SwapOptions struct {
}

type SwapOptionsUniversalRouter struct {
}
