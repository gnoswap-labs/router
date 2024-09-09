package math

import "math/big"

// GCD
// big.Int형 변수 2개를 입력으로 받아 두 수의 최대 공약수를 big.Int형으로 반환하는 함수이다.
func GCD(a, b *big.Int) *big.Int {
	zero := big.NewInt(0)

	aAbs := new(big.Int).Abs(a)
	bAbs := new(big.Int).Abs(b)

	for bAbs.Cmp(zero) != 0 {
		aAbs, bAbs = bAbs, new(big.Int).Mod(aAbs, bAbs)
	}
	return aAbs
}
