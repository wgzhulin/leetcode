package find_largest_value_in_each_tree_row

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhulinw/leetcode/utils"
	"testing"
)

func TestLargestValues(t *testing.T) {
	type data struct {
		input    []int
		excepted []int
	}

	testData := []data{
		{
			input:    []int{1, 3, 2, 5, 3, 0, 9},
			excepted: []int{1, 3, 9},
		},
		{
			input:    []int{1, 2, 3},
			excepted: []int{1, 3},
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.excepted, largestValues(utils.IntSliceToBinaryTree(tdata.input)))
		assert.Equal(t, tdata.excepted, largestValues2(utils.IntSliceToBinaryTree(tdata.input)))
	}
}
