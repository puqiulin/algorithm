package rotate_array

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
		input:    []int{1, 2, 3, 4, 5, 6, 7},
		input2:   3,
		expected: []int{5, 6, 7, 1, 2, 3, 4},
	},
	{
		name:     "Test2",
		input:    []int{-1, -100, 3, 99},
		input2:   2,
		expected: []int{3, 99, -1, -100},
	},
}

func TestRotateArray(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := RightRotateArray(test.input, test.input2)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}

func TestRotateArray2(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := RightRotateArray2(test.input, test.input2)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}

func TestRotateArray3(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := RightRotateArray3(test.input, test.input2)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}
