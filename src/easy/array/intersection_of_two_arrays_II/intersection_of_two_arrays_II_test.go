package intersect

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	name     string
	input    []int
	input2   []int
	expected [][]int
}

var tests = []Test{
	{
		name:     "Test1",
		input:    []int{1, 2, 2, 1},
		input2:   []int{2, 2},
		expected: [][]int{{2, 2}},
	},
	{
		name:     "Test2",
		input:    []int{4, 9, 5},
		input2:   []int{9, 4, 9, 8, 4},
		expected: [][]int{{4, 9}, {9, 4}},
	},
}

func TestIntersect(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := Intersect(test.input, test.input2)
			assert.Contains(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}

func TestIntersect2(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := Intersect2(test.input, test.input2)
			assert.Contains(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}
