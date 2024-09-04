package core

type AlphaRouterConfig struct {
	v3ProtocolPoolSelection ProtocolPoolSelection
	maxSwapsPerPath         int
	maxSplits               int
	minSplits               int
	distributionPercent     int
}
