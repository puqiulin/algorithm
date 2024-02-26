package rotate_image

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	name     string
	input    [][]int
	expected [][]int
}

var tests = []Test{
	{
		name:     "Test1",
		input:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		expected: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
	},
	{
		name:     "Test2",
		input:    [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
		expected: [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
	},
}

func TestRotate(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := rotate(test.input)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}
