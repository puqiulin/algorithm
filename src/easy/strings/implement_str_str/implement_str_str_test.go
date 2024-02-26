package implement_str_str

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	name     string
	input    string
	input2   string
	expected int
}

var tests = []Test{
	{
		name:     "Test1",
		input:    "sadbutsad",
		input2:   "sad",
		expected: 0,
	},
	{
		name:     "Test2",
		input:    "leetcode",
		input2:   "leeto",
		expected: -1,
	},
}

func TestIsAnagram(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := strStr(test.input, test.input2)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}
