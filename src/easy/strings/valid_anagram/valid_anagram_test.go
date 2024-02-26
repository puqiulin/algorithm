package valid_anagram

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	name     string
	input    string
	input2   string
	expected bool
}

var tests = []Test{
	{
		name:     "Test1",
		input:    "anagram",
		input2:   "nagaram",
		expected: true,
	},
	{
		name:     "Test2",
		input:    "rat",
		input2:   "car",
		expected: false,
	},
}

func TestIsAnagram(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := isAnagram(test.input, test.input2)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}

func TestIsAnagram2(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := isAnagram2(test.input, test.input2)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}

func TestIsAnagram3(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := isAnagram3(test.input, test.input2)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}
