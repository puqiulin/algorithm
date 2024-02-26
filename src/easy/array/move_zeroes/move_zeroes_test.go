package move_zeroes

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
		input:    []int{0, 1, 0, 3, 12},
		expected: []int{1, 3, 12, 0, 0},
	},
	{
		name:     "Test2",
		input:    []int{0},
		expected: []int{0},
	},
}

func TestMoveZeroes(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := moveZeroes(test.input)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}
