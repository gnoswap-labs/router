package math

import "math/big"

func GCD(a, b *big.Int) *big.Int {
	zero := big.NewInt(0)

	aAbs := new(big.Int).Abs(a)
	bAbs := new(big.Int).Abs(b)

	for bAbs.Cmp(zero) != 0 {
		aAbs, bAbs = bAbs, new(big.Int).Mod(aAbs, bAbs)
	}
	return aAbs
}
