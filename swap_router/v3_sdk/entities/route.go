package entities

import (
	"router/swap_router/core/entities/currency"
)

type Route[TInput currency.ICurrency, TOutput currency.ICurrency] struct {
}
