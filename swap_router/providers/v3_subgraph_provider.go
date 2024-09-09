package providers

import (
	"router/database"
	"router/swap_router"
	"router/swap_router/core/entities/currency"
)

const (
	defaultTrackedUSDThreshold = 0.01
	defaultLiquidityThreshold  = 100000
)

type IV3SubgraphProvider interface {
	getPools(tokenIn *currency.Token, tokenOut *currency.Token) []V3SubgraphPool
}

type V3SubgraphProvider struct {
	chainId             swap_router.ChainId
	db                  database.Database
	trackedUsdThreshold float64
	liquidityThreshold  float64
}

func NewV3SubgraphProvider(chainId swap_router.ChainId, db database.Database, trackedUSDThreshold, liquidityThreshold float64) *V3SubgraphProvider {
	if trackedUSDThreshold == 0.0 {
		trackedUSDThreshold = defaultTrackedUSDThreshold
	}
	if liquidityThreshold == 0.0 {
		liquidityThreshold = defaultLiquidityThreshold
	}

	return &V3SubgraphProvider{
		chainId:             chainId,
		db:                  db,
		trackedUsdThreshold: trackedUSDThreshold,
		liquidityThreshold:  liquidityThreshold,
	}
}

func (sp *V3SubgraphProvider) getPools(tokenIn *currency.Token, tokenOut *currency.Token) []V3SubgraphPool {
	pools := sp.db.GetPools()
	sanitizedPools := sp.selectPools(pools)
	return sanitizedPools
}

func (sp *V3SubgraphProvider) selectPools(pools []database.PoolV0) []V3SubgraphPool {

	return []V3SubgraphPool{}
}
