package swap_router

type AlphaRouterConfig struct {
	v3ProtocolPoolSelection ProtocolPoolSelection
	maxSwapsPerPath         int
	maxSplits               int
	minSplits               int

	/**
	 * The minimum percentage of the input token to use for each route in a split route.
	 * All routes will have a multiple of this value. For example is distribution percentage is 5,
	 * a potential return swap would be:
	 *
	 * 5% of input => Route 1
	 * 55% of input => Route 2
	 * 40% of input => Route 3
	 */
	distributionPercent int

	//
	useCachedRoutes        bool
	writeToCachedRoutes    bool
	optimisticCachedRoutes bool
}
