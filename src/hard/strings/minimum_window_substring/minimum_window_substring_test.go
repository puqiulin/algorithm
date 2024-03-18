package minimum_window_substring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	name     string
	input    string
	input2   string
	expected string
}

var tests = []Test{
	{
		name:     "Test1",
		input:    "ADOBECODEBANC",
		input2:   "ABC",
		expected: "BANC",
	},
	{
		name:     "Test2",
		input:    "a",
		input2:   "a",
		expected: "a",
	},
	{
		name:     "Test2",
		input:    "a",
		input2:   "aa",
		expected: "",
	},
}

func TestStrStr(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := minWindow(test.input, test.input2)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}
