package problems

import (
	"github.com/zhulinw/leetcode/testify/assert"
	"github.com/zhulinw/leetcode/testify/tutils"
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
		assert.Equal(t, tutils.SliceToListNode(tdata.excepted), swapPairs(tutils.SliceToListNode(tdata.input)))
		assert.Equal(t, tutils.SliceToListNode(tdata.excepted), swapPairs2(tutils.SliceToListNode(tdata.input)))
	}
}
