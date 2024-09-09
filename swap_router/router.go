package swap_router

import (
	currency2 "router/swap_router/core/entities/currency"
	fractions2 "router/swap_router/core/entities/fractions"
	"router/swap_router/core/entities/fractions/math"
	"router/swap_router/core/fractions"
	entities2 "router/swap_router/router-sdk/entities"
	"router/swap_router/v3_sdk/entities"
)

type V3Route struct {
	entities.Route[currency2.Token, currency2.Token]
	Protocol entities2.Protocol
}

type SwapRoute struct {
	quote         fractions2.CurrencyAmount
	trade         entities2.Trade[currency2.ICurrency, currency2.ICurrency, TradeType]
	route         []V3RouteWithValidQuote
	portionAmount fractions2.CurrencyAmount
}

type SwapToRatioRoute struct {
	SwapRoute
	optimalRatio       math.Fraction
	postSwapTargetPool Pool
}

type SwapToRatioStatus int

const (
	SUCCESS        SwapToRatioStatus = 1
	NO_ROUTE_FOUND SwapToRatioStatus = 2
	NO_SWAP_NEEDED SwapToRatioStatus = 3
)

type SwapToRatioSuccess struct {
	Status SwapToRatioStatus
	Result SwapToRatioRoute
}

func (s *SwapToRatioSuccess) GetStatus() SwapToRatioStatus {
	return s.Status
}

type SwapToRatioFailure struct {
	Status SwapToRatioStatus
	Error  string
}

func (s *SwapToRatioFailure) GetStatus() SwapToRatioStatus {
	return s.Status
}

type SwapToRatioNoSwapNeeded struct {
	Status SwapToRatioStatus
}

func (s *SwapToRatioNoSwapNeeded) GetStatus() SwapToRatioStatus {
	return s.Status
}

type SwapToRatioResponse interface {
	GetStatus() SwapToRatioStatus
}

type SwapType int

const (
	UNIVERSAL_ROUTER SwapType = iota
	SWAP_ROUTER_O2
)

type FlatFeeOptions struct {
	Amount    BigintIsh // 금액 필드
	Recipient string    // 수신자 필드
}

type RouterSwapOptions struct {
	flatFee           FlatFeeOptions
	slippageTolerance fractions.Percent
	recipient         string
	fee               FeeOptions
}

type FeeOptions struct {
	fee       fractions.Percent
	recipient string
}

type SwapOptionsUniversalRouter struct {
	RouterSwapOptions
	Type     SwapType
	Simulate *Simulate
}

type Simulate struct {
	FromAddress string
}

type SwapOptionsSwapRouter02 struct {
	Type              SwapType          // SwapType 필드
	Recipient         string            // 수신자 주소
	SlippageTolerance fractions.Percent // 슬리피지 허용치
	Deadline          int64             // 마감 시간 (Unix 타임스탬프)
	Simulate          *Simulate         // 선택적 필드
	InputTokenPermit  *InputTokenPermit // 선택적 필드
}

type InputTokenPermit struct {
	V        int     // 0, 1, 27, 28 중 하나
	R        string  // 문자열 타입
	S        string  // 문자열 타입
	Amount   *string // 선택적 필드
	Deadline *string // 선택적 필드
	Nonce    *string // 선택적 필드
	Expiry   *string // 선택적 필드
}

// 원문: type SwapOptions = SwapOptionsUniversalRouter | SwapOptionsSwapRouter02
type SwapOptions struct {
	UniversalRouter *SwapOptionsUniversalRouter // 선택적 필드
	SwapRouter02    *SwapOptionsSwapRouter02    // 선택적 필드
}

type IRouter[RoutingConfig any] interface {
	Route(amount fractions2.CurrencyAmount, quoteCurrency currency2.ICurrency, swapType TradeType, swapOptions *SwapOptions, partialRoutingConfig RoutingConfig) *SwapRoute
}
