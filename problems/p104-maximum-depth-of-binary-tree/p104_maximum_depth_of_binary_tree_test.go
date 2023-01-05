package maximum_depth_of_binary_tree

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhulinw/leetcode/utils"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	type Input struct {
		Para1 []int
	}

	type Output struct {
		Para1 int
	}

	input := []Input{
		{
			Para1: []int{3, 9, 20, 0, 0, 15, 7},
		},
		{
			Para1: []int{1, 0, 2},
		},
	}

	output := []Output{
		{Para1: 3},
		{Para1: 2},
	}

	for i := range input {
		assert.Equal(t, output[i].Para1, maxDepth(utils.IntSliceToBinaryTree(input[i].Para1)))
	}
}
