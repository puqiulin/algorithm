package string_to_integer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	name     string
	input    string
	expected int
}

var tests = []Test{
	{
		name:     "Test1",
		input:    "42",
		expected: 42,
	},
	{
		name:     "Test2",
		input:    "   -42",
		expected: -42,
	},
	{
		name:     "Test3",
		input:    "4193 with words",
		expected: 4193,
	},
}

func TestIsAnagram(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := myAtoi(test.input)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}
