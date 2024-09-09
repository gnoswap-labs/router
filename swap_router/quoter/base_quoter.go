package quoter

import "router/swap_router"

type BaseQuoter struct {
	chainId  swap_router.ChainId
	protocol Protocol
}
