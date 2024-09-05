package quoter

import "router/core"

type BaseQuoter struct {
	chainId  core.ChainId
	protocol Protocol
}
