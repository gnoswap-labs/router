package core

type TradeType int

const (
	EXACT_INPUT = iota
	EXACT_OUTPUT
)

type FeeAmount int

const (
	LOWEST = 100
	LOW    = 500
	MEDIUM = 3000
	HIGH   = 10000
)
