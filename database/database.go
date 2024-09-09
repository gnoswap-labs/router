package database

import "context"

type Database interface {
	GetPools() []PoolV0
	GetTokens() []TokenV0
	Query(ctx context.Context, sql string) (interface{}, error) // 제너릭 타입은 Go에서 다루기 복잡하므로 interface{} 사용
	Commit(ctx context.Context, sql string, data [][]interface{}) (bool, error)
}
