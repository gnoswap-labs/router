package poc

// Pool
// kind of edge
type Pool struct {
	Address string

	TokenA Token // 효율을 위해서는 포인터로 가져오는게 좋을만 하다.
	TokenB Token

	ReserveA float64
	ReserveB float64

	// 거래 수수료
	//Fee uint32
}
