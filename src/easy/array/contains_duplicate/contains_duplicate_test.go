package contains_duplicate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	name     string
	input    []int
	expected bool
}

var tests = []Test{
	{
		name:     "Test1",
		input:    []int{1, 2, 3, 1},
		expected: true,
	},
	{
		name:     "Test2",
		input:    []int{1, 2, 3, 4},
		expected: false,
	},
	{
		name:     "Test2",
		input:    []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
		expected: true,
	},
}

func TestContainsDuplicate(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := containsDuplicate(test.input)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}

func TestContainsDuplicate2(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := containsDuplicate2(test.input)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}

func TestContainsDuplicate3(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := containsDuplicate3(test.input)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}
