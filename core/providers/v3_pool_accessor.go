package providers

import (
	"router/core"
	"router/core/currency"
)

type V3PoolAccessor interface {
	GetPool(tokenA currency.Token, tokenB currency.Token, feeAmount core.FeeAmount) (*core.Pool, error) // Pool 또는 오류 반환
	GetPoolByAddress(address string) (*core.Pool, error)                                                // Pool 또는 오류 반환
	GetAllPools() []core.Pool                                                                           // Pool 배열 반환
}
