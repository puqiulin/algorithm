package first_unique_char

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
		input:    "leetcode",
		expected: 0,
	},
	{
		name:     "Test2",
		input:    "loveleetcode",
		expected: 2,
	},
	{
		name:     "Test3",
		input:    "aabb",
		expected: -1,
	},
}

func TestFirstUniqChar(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := firstUniqChar(test.input)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}
