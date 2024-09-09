package providers

import (
	"router/swap_router/core/entities/currency"
)

type V3SubgraphPool struct {
	ID                  string         `json:"id"`
	Liquidity           string         `json:"liquidity"`
	TokenA              currency.Token `json:"tokenA"`
	TokenB              currency.Token `json:"tokenB"`
	TotalValueLockedUSD float64        `json:"tvlUSD"`
	Fee                 float64        `json:"fee"`
}

type RawV3SubgraphPool struct {
	ID                  string
	Liquidity           string
	TokenAPath          string
	TokenBPath          string
	TotalValueLockedUSD string
	Fee                 float64
}
