package providers

import (
	"router/database"
	"router/swap_router"
	"router/swap_router/core/entities/currency"
)

// pool_privider.getPools에서 사용되는 구조체
// getPools에서 인자를 적절하게 넘겨만 받으면 삭제해도 된다.
// 이름만 적절하게 지으면 잘 활용할 수 있을 거 같은데
type TokenPairs struct {
	tokenA    currency.Token
	tokenB    currency.Token
	feeAmount swap_router.FeeAmount
}

type IV3PoolProvider interface {
	getPools(tokenPairs []TokenPairs) IV3PoolAccessor
	getPoolAddress(tokenA, tokenB currency.Token, feeAmount swap_router.FeeAmount) (string, currency.Token, currency.Token)
}

type V3PoolProvider struct {
	PoolAddressCache map[string]string
	chainId          swap_router.ChainId
	database         database.Database
}

func NewV3PoolProvider(chainId swap_router.ChainId, db database.Database) *V3PoolProvider {
	return &V3PoolProvider{
		PoolAddressCache: make(map[string]string),
		chainId:          chainId,
		database:         db,
	}
}

// todo: goroutine
func (pp *V3PoolProvider) getPools(tokenPairs []TokenPairs) IV3PoolAccessor {
	poolAddressSet := make(map[string]struct{})
	sortedTokenPairs := make([]TokenPairs, 0)
	sortedPoolAddresses := make([]string, 0)

	for _, tokenPair := range tokenPairs {
		poolAddress, token0, token1 := pp.getPoolAddress(tokenPair.tokenA, tokenPair.tokenB, tokenPair.feeAmount)

		if _, exists := poolAddressSet[poolAddress]; exists {
			continue
		}

		poolAddressSet[poolAddress] = struct{}{}
		sortedTokenPairs = append(sortedTokenPairs, tokenPair)
		sortedPoolAddresses = append(sortedPoolAddresses, poolAddress)
	}

	//slot0Results, liquidityResults

	poolAddressToPool := make(map[string]struct{})
	invalidPools := make(map[string]struct{})

	for i, _ := range sortedPoolAddresses {
		slot0Result := slot0Results[i]
		liquidityResult := liquidityResults[i]

		token0, token1, fee := sortedPoolAddresses[i]

		slot0 := slot0Result
	}
}

func (pp *V3PoolProvider) getPoolAddress(tokenA, tokenB currency.Token, feeAmount swap_router.FeeAmount) (string, currency.Token, currency.Token) {
	token0, token1 := tokenA.SortsByLowerAddress(tokenB)
	//cacheKey =
}

func (pp **V3PoolProvider) getPoolsData(poolAddresses string, dataName string) {
	pools := pp.database.GetPools()
	//results := make()

	for _, address := range poolAddresses {
		//pool := pools.f
	}
}
