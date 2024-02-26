package reverse_string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	name     string
	input    int
	expected int
}

var tests = []Test{
	{
		name:     "Test1",
		input:    123,
		expected: 321,
	},
	{
		name:     "Test2",
		input:    -123,
		expected: -321,
	},
	{
		name:     "Test3",
		input:    120,
		expected: 21,
	},
}

func TestReverseString(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := reverseInteger(test.input)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}
