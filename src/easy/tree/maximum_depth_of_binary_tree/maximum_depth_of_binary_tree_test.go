package implement_str_str

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Test struct {
	name     string
	input    []TreeNode
	expected int
}

var input1 []TreeNode
s:=[]int{3, 9, 20, 0, 0, 15, 7}
for _,v :=range s{

}

var tests = []Test{
	{
		name:     "Test1",
		input:    input1,
		expected: 3,
	},
	{
		name:     "Test2",
		input:    []int{1, nil, 2},
		expected: 2,
	},
}

func TestMaxDepth(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := maxDepth(test.input)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}

func TestMaxDepth2(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := maxDepth2(test.input)
			assert.Equal(t, test.expected, output, fmt.Sprintf("error: %v vs %v", output, test.expected))
		})
	}
}
