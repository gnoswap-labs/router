package poc

import (
	"fmt"
	"math"
	"testing"
)

const tolerance = 0.00000001 // 오차 범위

func TestMyRouterV1(t *testing.T) {
	tokens := map[string]Token{
		"a": {Symbol: "a"},
		"b": {Symbol: "b"},
	}

	tests := []struct {
		name     string
		edges    []*Pool
		requests []SwapRequest
		results  []SwapResult
	}{
		{
			name: "단일 홉 스왑",
			edges: []*Pool{
				{"a:b", tokens["a"], tokens["b"], 4000, 1000},
			},
			requests: []SwapRequest{
				{"a", "b", 2000},
			},
			results: []SwapResult{
				{"a", "b", 2000.0, 2000.0 / 6.0},
			},
		},
		// TODO: 아래 테스트 케이스도 통과해야 함. 값은 검증 필요. 
		// {
		// 	name: "극단적인 비율 스왑",
		// 	edges: []*Pool{
		// 		{"a:b", tokens["a"], tokens["b"], 1000000, 1},
		// 	},
		// 	requests: []SwapRequest{
		// 		{"a", "b", 500},
		// 	},
		// 	results: []SwapResult{
		// 		{"a", "b", 500, 0.0004999999999999999},
		// 	},
		// },
		// {
		// 	name: "양방향 스왑",
		// 	edges: []*Pool{
		// 		{"a:b", tokens["a"], tokens["b"], 4000, 1000},
		// 		{"b:a", tokens["b"], tokens["a"], 1000, 4000},
		// 	},
		// 	requests: []SwapRequest{
		// 		{"a", "b", 2000},
		// 		{"b", "a", 500},
		// 	},
		// 	results: []SwapResult{
		// 		{"a", "b", 2000, 2000.0 / 6.0},
		// 		{"b", "a", 500, 500.0 / 6.0},
		// 	},
		// },
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
	tokens := map[string]Token{
		"a": {Symbol: "a"},
		"b": {Symbol: "b"},
		"c": {Symbol: "c"},
		"d": {Symbol: "d"},
	}

	tests := []struct {
		name            string
		edges           []*Pool
		requests        []SwapRequest
		results         []SwapResult
		maxSearchLength int
		expectError     bool
	}{
		{
			name: "최대 검색 길이 1의 다중 홉 스왑",
			edges: []*Pool{
				{"a:b", tokens["a"], tokens["b"], 4000, 1000},
				{"a:c", tokens["a"], tokens["c"], 2000, 1000},
				{"b:c", tokens["b"], tokens["c"], 2000, 4000}},
			requests: []SwapRequest{
				{"a", "c", 2000}},
			results: []SwapResult{
				{"a", "c", 2000, 571.4285714285},
			},
			maxSearchLength: 2,
		},
		{
			name: "최대 검색 길이 2의 다중 홉 스왑",
			edges: []*Pool{
				{"a:b", tokens["a"], tokens["b"], 4000, 1000},
				{"a:c", tokens["a"], tokens["c"], 2000, 1000},
				{"b:c", tokens["b"], tokens["c"], 2000, 4000},
			},
			requests: []SwapRequest{
				{"a", "c", 2000},
			},
			results: []SwapResult{
				{"a", "c", 2000, 500},
			},
			maxSearchLength: 1,
		},
		{
			name: "다양한 경로 스왑",
			edges: []*Pool{
				{"a:b", tokens["a"], tokens["b"], 4000, 1000},
				{"b:c", tokens["b"], tokens["c"], 2000, 4000},
				{"a:c", tokens["a"], tokens["c"], 2000, 1000},
				{"c:d", tokens["c"], tokens["d"], 1000, 500},
			},
			requests: []SwapRequest{
				{"a", "d", 1000},
			},
			results: []SwapResult{
				{"a", "d", 1000,  133.33333333333334}, // TODO: 임의로 넣은 값이라 검증 필요
			},
			maxSearchLength: 3,
		},	
		{
			name: "검색 길이가 음수인 경우",
			edges: []*Pool{
				{"a:b", tokens["a"], tokens["b"], 4000, 1000},
				{"b:c", tokens["b"], tokens["c"], 2000, 4000},
				{"c:d", tokens["c"], tokens["d"], 1000, 500},
				{"a:d", tokens["a"], tokens["d"], 3000, 1500},
			},
			requests: []SwapRequest{
				{"a", "d", 1500},
			},
			maxSearchLength: -1,
			expectError:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			router := NewMyRouter(test.edges)

			for i, request := range test.requests {
				result, err := router.findRouteV2(request, test.maxSearchLength, nil)
				if err != nil {
					if !test.expectError {
						t.Fatalf("Router: can't find path: %v:%v", request.FromTokenSymbol, request.ToTokenSymbol)
					}
					continue
				}

				diff := math.Abs(result[len(result)-1].AmountOut - test.results[i].AmountOut)
				if diff > tolerance {
					t.Fatalf("Router: Unexpected Token output number, expected: %v, got %v", test.results[i].AmountOut, result[len(result)-1].AmountOut)
				}
			}
		})
	}
}
