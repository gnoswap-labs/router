package swap_router

type ProtocolPoolSelection struct {
	topN                  int
	topNDirectSwaps       int
	topNTokenInOut        int
	topNSecondHop         int
	topNWithEachBaseToken int
	topNWithBaseToken     int

	// selectable variable
	topNSecondHopForTokenAddress map[string]int
	tokensToAvoidOnSecondHops    []string
}
