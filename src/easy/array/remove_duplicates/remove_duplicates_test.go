package remove_duplicates

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
		input:    []int{1, 1, 2},
		expected: 2,
	},
	{
		name:     "Test2",
		input:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		expected: 5,
	},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := RemoveDuplicates(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}

func TestRemoveDuplicates2(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := RemoveDuplicates2(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}
