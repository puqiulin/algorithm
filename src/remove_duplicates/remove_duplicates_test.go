package remove_duplicates

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Test1",
			input:    []int{1, 1, 2},
			expected: 2,
		},
		{
			name:     "Test2",
			input:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := RemoveDuplicates(test.input)
			assert.Equal(t, output, test.expected, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}

func TestRemoveDuplicates2(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Test1",
			input:    []int{1, 1, 2},
			expected: 2,
		},
		{
			name:     "Test2",
			input:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := RemoveDuplicates2(test.input)
			assert.Equal(t, output, test.expected, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}
