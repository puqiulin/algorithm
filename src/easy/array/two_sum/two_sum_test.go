package two_sum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	name     string
	input    []int
	input2   int
	expected []int
}

var tests = []Test{
	{
		name:     "Test1",
		input:    []int{2, 7, 11, 15},
		input2:   9,
		expected: []int{0, 1},
	},
	{
		name:     "Test2",
		input:    []int{3, 2, 4},
		input2:   6,
		expected: []int{1, 2},
	},
	{
		name:     "Test3",
		input:    []int{3, 3},
		input2:   6,
		expected: []int{0, 1},
	},
}

func TestTwoSum(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := twoSum(test.input, test.input2)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}

func TestTwoSum2(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := twoSum2(test.input, test.input2)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}
