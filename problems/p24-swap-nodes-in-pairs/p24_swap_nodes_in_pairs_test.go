package swap_nodes_in_pairs

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhulinw/leetcode/utils"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	type data struct {
		input    []int
		excepted []int
	}

	testData := []data{
		{
			input:    []int{1, 2, 3, 4},
			excepted: []int{2, 1, 4, 3},
		},
		{
			input:    []int{},
			excepted: []int{},
		},
		{
			input:    []int{1},
			excepted: []int{1},
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, utils.SliceToListNode(tdata.excepted), swapPairs(utils.SliceToListNode(tdata.input)))
		assert.Equal(t, utils.SliceToListNode(tdata.excepted), swapPairs2(utils.SliceToListNode(tdata.input)))
	}
}
