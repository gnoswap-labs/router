package core

import "router/core/providers"

type AlphaRouterParams struct {
	// The chain id for this instance of the Alpha Router.
	chainId ChainId
	// The provider for getting all pools that exist on V3 from the Subgraph. The pools
	// from this provider are filtered during the algorithm to a set of candidate pools.
	v3SubgraphProvider *providers.IV3SubgraphProvider
	// The provider for getting data about V3 pools.
	v3PoolProvider *providers.IV3PoolProvider
	// The provider for getting V3 quotes.
	onChainQuoteProvider *IOnChainQuoteProvider
	// The provider for getting data about Tokens.
	tokenProvider *ITokenProvider
	// A provider for computing the portion-related data for routes and quotes.
	portionProvider *IPortionProvider
}
