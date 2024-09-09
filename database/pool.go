package database

// V0이 붙네
type PoolV0 struct {
	id     string
	height int
	//time Date
	pool_path               string
	token0_path             string
	token1_path             string
	token0_balance          string
	token1_balance          string
	fee                     int
	tick_spacing            int
	max_liquidity_per_tick  string
	sqrt_price_x96          string
	tick                    int
	fee_protocol            int
	unlocked                bool
	fee_growth_global0_x128 string
	fee_growth_global1_x128 string
	protocol_fee_token0     string
	protocol_fee_token1     string
	liquidity               string
	ticks                   map[string]PoolTickV0 // ticks: { [key in string]: PoolTickVO };
	tick_bitmaps            map[string]string     // tick_bitmaps: { [key in string]: string };
	positions               []PoolPositionV0
	tvl_usd                 string
}

type PoolTickV0 struct {
	initialized                 bool
	liquidityNet                string
	liquidityGross              string
	secondsOutside              int
	feeGrowthOutside0X128       string // 애는 갑자기 OX128이야 0_x128이 아니라
	feeGrowthOutside1X128       string
	tickCumulativeOutside       int
	secondsPerLiquidityOutsideX string
}

type PoolPositionV0 struct {
	owner      string
	liquidity  string
	tickLower  int
	tickUpper  int
	token0Owed string
	token1Owed string
}
