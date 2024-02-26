package single_number

import (
	"fmt"
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
		input:    []int{2, 2, 1},
		expected: 1,
	},
	{
		name:     "Test2",
		input:    []int{4, 1, 2, 1, 2},
		expected: 4,
	},
	{
		name:     "Test3",
		input:    []int{1},
		expected: 1,
	},
}

func TestSingleNumber(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := SingleNumber(test.input)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}

func TestSingleNumber2(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := SingleNumber2(test.input)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}

func TestSingleNumber3(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := SingleNumber3(test.input)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}

func TestSingleNumber4(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := SingleNumber4(test.input)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}
