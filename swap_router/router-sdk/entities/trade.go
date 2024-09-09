package entities

import (
	"router/swap_router"
	"router/swap_router/core/entities/currency"
)

type Trade[TInput currency.ICurrency, TOutput currency.ICurrency, TTradeType swap_router.TradeType] struct {
}
