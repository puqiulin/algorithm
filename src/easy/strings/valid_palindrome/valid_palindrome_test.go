package valid_palindrome

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	name     string
	input    string
	expected bool
}

var tests = []Test{
	{
		name:     "Test1",
		input:    "A man, a plan, a canal: Panama",
		expected: true,
	},
	{
		name:     "Test2",
		input:    "race a car",
		expected: false,
	},
	{
		name:     "Test3",
		input:    " ",
		expected: true,
	},
}

func TestIsAnagram(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := isPalindrome(test.input)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}
