package poc

import (
	"fmt"
	"math"
	"testing"
)

func TestMyRouterV1(t *testing.T) {
	tolerance := 0.00000001 // 오차 범위
	tokens := map[string]Token{
		"a": Token{Symbol: "a"},
		"b": Token{Symbol: "b"},
	}

	tests := []struct {
		edges    []*Pool
		requests []SwapRequest
		results  []SwapResult
	}{
		{
			[]*Pool{
				{"a:b", tokens["a"], tokens["b"], 4000, 1000}},
			[]SwapRequest{
				{"a", "b", 2000}},
			[]SwapResult{
				{"a", "b", 2000.0, 2000.0 / 6.0},
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			router := NewMyRouter(test.edges)

			for i, request := range test.requests {
				result, err := router.findRouteV1(request)
				if err != nil {
					t.Fatalf("Swap Error: can't find pool: %v:%v", request.FromTokenSymbol, request.ToTokenSymbol)
				}

				diff := math.Abs(result[0].AmountOut - test.results[i].AmountOut)
				if diff > tolerance {
					t.Fatalf("Swap: Unexpected Token output number, expected: %v, got %v", test.results[i].AmountOut, result[0].AmountOut)
				}
				fmt.Println(result[0])

				fmt.Println("스왑 결과")
				for _, pool := range router.network {
					fmt.Printf("pool (%s) %s: %f %s: %f\n", pool.Address, pool.TokenA.Symbol, pool.ReserveA, pool.TokenB.Symbol, pool.ReserveB)
				}
			}
		})
	}
}

func TestMyRouterV2(t *testing.T) {
	tolerance := 0.00000001 // 오차 범위
	tokens := map[string]Token{
		"a": Token{Symbol: "a"},
		"b": Token{Symbol: "b"},
		"c": Token{Symbol: "c"},
		"d": Token{Symbol: "d"},
	}

	tests := []struct {
		name            string
		edges           []*Pool
		requests        []SwapRequest
		results         []SwapResult
		maxSearchLength int
	}{
		{
			"최대 검색 길이 1의 다중 홉 스왑",
			[]*Pool{
				{"a:b", tokens["a"], tokens["b"], 4000, 1000},
				{"a:c", tokens["a"], tokens["c"], 2000, 1000},
				{"b:c", tokens["b"], tokens["c"], 2000, 4000}},
			[]SwapRequest{
				{"a", "c", 2000}},
			[]SwapResult{
				{"a", "c", 2000, 571.4285714285},
			},
			2,
		},
		{
			"최대 검색 길이 2의 다중 홉 스왑",
			[]*Pool{
				{"a:b", tokens["a"], tokens["b"], 4000, 1000},
				{"a:c", tokens["a"], tokens["c"], 2000, 1000},
				{"b:c", tokens["b"], tokens["c"], 2000, 4000}},
			[]SwapRequest{
				{"a", "c", 2000}},
			[]SwapResult{
				{"a", "c", 2000, 500},
			},
			1,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			router := NewMyRouter(test.edges)

			for i, request := range test.requests {
				result, err := router.findRouteV2(request, test.maxSearchLength, nil)
				if err != nil {
					t.Fatalf("Router: can't find path: %v:%v", request.FromTokenSymbol, request.ToTokenSymbol)
				}
				fmt.Print("result path: ")
				fmt.Println(result)

				diff := math.Abs(result[len(result)-1].AmountOut - test.results[i].AmountOut)
				if diff > tolerance {
					t.Fatalf("Router: Unexpected Token output number, expected: %v, got %v", test.results[i].AmountOut, result[len(result)-1].AmountOut)
				}
			}
		})
	}
}
