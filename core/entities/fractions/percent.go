package fractions

import "router/core/math"

type Percent struct {
	math.Fraction
}

func ToPercent(f math.Fraction) Percent {
	return Percent{
		Fraction: f,
	}
}
