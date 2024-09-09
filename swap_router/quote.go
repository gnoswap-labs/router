package swap_router

import (
	"router/swap_router/core/entities/fractions/math"
)

type Quote struct {
	percent *math.Fraction
	amount  *math.Fraction
}
