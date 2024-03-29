package merge_two_sorted_lists

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhulinw/leetcode/utils"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	type data struct {
		input1   []int
		input2   []int
		excepted []int
	}

	testData := []data{
		{
			input1:   []int{1, 2, 4},
			input2:   []int{1, 3, 4},
			excepted: []int{1, 1, 2, 3, 4, 4},
		},
		{
			input1:   []int{},
			input2:   []int{},
			excepted: []int{},
		},
		{
			input1:   []int{},
			input2:   []int{0},
			excepted: []int{0},
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, utils.SliceToListNode(tdata.excepted),
			mergeTwoLists(utils.SliceToListNode(tdata.input1), utils.SliceToListNode(tdata.input2)))
	}
}
