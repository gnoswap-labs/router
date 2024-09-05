package core

import (
	"router/core/currency"
	"router/core/providers"
)

// 추측하기로 사용 가능한 풀들을 땡겨오는 무언가?

type V3CandidatePools struct {
	poolAccessor   providers.V3PoolAccessor
	candidatePools CandidatePoolsBySelectionCriteria
	subgraphPools  []providers.V3SubgraphPool
}

func getV3CandidatePools(
	tokenIn currency.Token,
	tokenOut currency.Token,
	tradeType TradeType,
	routerConfig AlphaRouterConfig,
	//subgraphProvider,
	//tokenProvider,
	//poolPrivider,
	chainId ChainId,
) V3CandidatePools {
	// ????? Protocol이 왜 들어가
	v3ProtocolPoolSelection := routerConfig.v3ProtocolPoolSelection

	tokenInAddress := tokenIn.Address
	tokenOutAddress := tokenOut.Address

	//allPools :=

}
