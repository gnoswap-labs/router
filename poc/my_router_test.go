package poc

import (
	"fmt"
	"math"
	"testing"
)

func TestMyRouter(t *testing.T) {
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
				{2000.0, 2000.0 / 6.0},
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			router := NewMyRouter(test.edges)

			for i, request := range test.requests {
				result, err := router.Swap(request)
				if err != nil {
					t.Fatalf("Swap Error: can't find pool: %v:%v", request.FromToken, request.ToToken)
				}

				diff := math.Abs(result.AmountOut - test.results[i].AmountOut)
				tolerance := 0.00000001
				if diff > tolerance {
					t.Fatalf("Swap: Unexpected Token output number, expected: %v, got %v", test.results[i].AmountOut, result.AmountOut)
				}
				fmt.Println(result)

				for _, pool := range router.network {
					fmt.Println(pool)
				}
			}
		})
	}
}
