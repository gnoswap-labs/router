package core

type ProtocolPoolSelection struct {
	topN                  int
	topNDirectSwaps       int
	topNTokenInOut        int
	topNSecondHop         int
	topNWithEachBaseToken int
	topNWithBaseToken     int

	// selectable variable
	//topNSecondHopForTokenAddress
	//tokensToAvoidOnSecondHops
}
