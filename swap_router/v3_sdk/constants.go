package v3_sdk

const ADDRESS_ZERO = "0x0000000000000000000000000000000000000000"

// The default factory enabled fee amounts, denominated in hundredths of bips.
type FeeAmount int

const (
	LOWEST = 100
	LOW    = 500
	MEDIUM = 3000
	HIGH   = 10000
)

// The default factory tick spacings by fee amount.
var TICK_SPACINGS = map[FeeAmount]int{
	LOWEST: 1,
	LOW:    10,
	MEDIUM: 60,
	HIGH:   200,
}
