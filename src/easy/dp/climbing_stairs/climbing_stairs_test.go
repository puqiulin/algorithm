package climbing_stairs

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
		input:    2,
		expected: 2,
	},
	{
		name:     "Test2",
		input:    3,
		expected: 3,
	},
}

func TestStrStr(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := climbStairs(test.input)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}
