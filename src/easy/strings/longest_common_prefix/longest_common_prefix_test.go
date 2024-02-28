package longest_common_prefix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	name     string
	input    []string
	expected string
}

var tests = []Test{
	{
		name:     "Test1",
		input:    []string{"flower", "flow", "flight"},
		expected: "fl",
	},
	{
		name:     "Test2",
		input:    []string{"dog", "racecar", "car"},
		expected: "",
	},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := longestCommonPrefix(test.input)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}
