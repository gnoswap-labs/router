package providers

import (
	"router/core"
	"router/core/currency"
	"router/database"
)

type IV3PoolProvider interface {
	getPools(tokenPairs ??) V3PoolAccessor
	getPoolAddress(tokenA, tokenB currency.Token, feeAmount core.FeeAmount) (string, currency.Token, currency.Token)
}

type V3PoolProvider struct {
	PoolAddressCache map[string]string
	database database.Database
}

func NewV3PoolProvider(chainId core.ChainId, db database.Database) *V3PoolProvider {
	return &V3PoolProvider{
		PoolAddressCache: make(map[string]string),
		database:         db,
	}
}


// todo: goroutine
func (pp *V3PoolProvider) getPools() {

}

func (pp *V3PoolProvider) getPoolAddress(tokenA, tokenB currency.Token, feeAmount core.FeeAmount) (string, currency.Token, currency.Token) {

}