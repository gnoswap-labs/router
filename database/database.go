package database

type Database interface {
	GetPools() []PoolV0
	GetTokens() []TokenV0
	//query(sql string) T ????
	//commit(string, any[][]) bool ??????
}
