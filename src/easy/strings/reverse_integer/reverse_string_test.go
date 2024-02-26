package reverse_integer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	name     string
	input    []string
	expected []string
}

var tests = []Test{
	{
		name:     "Test1",
		input:    []string{"h", "e", "l", "l", "o"},
		expected: []string{"o", "l", "l", "e", "h"},
	},
	{
		name:     "Test2",
		input:    []string{"H", "a", "n", "n", "a", "h"},
		expected: []string{"h", "a", "n", "n", "a", "H"},
	},
}

func TestReverseString(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := reverseString(test.input)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}

func TestReverseString2(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := reverseString2(test.input)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}
