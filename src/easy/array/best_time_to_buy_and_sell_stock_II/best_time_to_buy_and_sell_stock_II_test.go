package best_time_to_buy_and_sell_stock_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	name     string
	input    []int
	expected int
}

var tests = []Test{
	{
		name:     "Test1",
		input:    []int{7, 1, 5, 3, 6, 4},
		expected: 7,
	},
	{
		name:     "Test2",
		input:    []int{1, 2, 3, 4, 5},
		expected: 4,
	},
	{
		name:     "Test3",
		input:    []int{7, 6, 4, 3, 1},
		expected: 0,
	},
}

func TestBestTimeToBuyAndSellStockII(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := BestTimeToBuyAndSellStockII(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}
