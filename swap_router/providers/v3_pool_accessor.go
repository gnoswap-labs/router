package providers

import (
	"router/swap_router"
	"router/swap_router/core/entities/currency"
)

type IV3PoolAccessor interface {
	GetPool(tokenA currency.Token, tokenB currency.Token, feeAmount swap_router.FeeAmount) (*swap_router.Pool, error) // Pool 또는 오류 반환
	GetPoolByAddress(address string) (*swap_router.Pool, error)                                                       // Pool 또는 오류 반환
	GetAllPools() []swap_router.Pool                                                                                  // Pool 배열 반환
}
