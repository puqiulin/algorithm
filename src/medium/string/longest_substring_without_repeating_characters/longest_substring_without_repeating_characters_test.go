package longest_substring_without_repeating_characters

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
		input:    "abcabcbb",
		expected: 3,
	},
	{
		name:     "Test2",
		input:    "bbbbb",
		expected: 1,
	},
	{
		name:     "Test3",
		input:    "pwwkew",
		expected: 3,
	},
}

func TestStrStr(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := lengthOfLongestSubstring(test.input)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}
