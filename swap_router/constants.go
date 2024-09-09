package swap_router

type BigintIsh interface {
}

type TradeType int

const (
	EXACT_INPUT = iota
	EXACT_OUTPUT
)

var MIN_PRICE_X96 = 4295128740
var MAX_PRICE_X96 = 1461446703485210103287273052203988822378723970341

const MIN_TICK = -887272
const MAX_TICK = 887272

const DEFAULT_SWAP_FEE = 0.15
