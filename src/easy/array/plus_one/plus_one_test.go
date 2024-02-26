package plus_one

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	name     string
	input    []int
	expected []int
}

var tests = []Test{
	{
		name:     "Test1",
		input:    []int{1, 2, 3},
		expected: []int{1, 2, 4},
	},
	{
		name:     "Test2",
		input:    []int{4, 3, 2, 1},
		expected: []int{4, 3, 2, 2},
	},
	{
		name:     "Test3",
		input:    []int{9},
		expected: []int{1, 0},
	},
}

func TestPlusOne(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := plusOne(test.input)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}
