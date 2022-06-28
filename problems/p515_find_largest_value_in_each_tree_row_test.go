package problems

import (
	"github.com/zhulinw/leetcode/testify/assert"
	"github.com/zhulinw/leetcode/testify/tutils"
	"testing"
)

/*
示例1：
输入: root = [1,3,2,5,3,null,9]
输出: [1,3,9]

示例2：
输入: root = [1,2,3]
输出: [1,3]
*/
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
		assert.Equal(t, tdata.excepted, largestValues(tutils.IntSliceToBinaryTree(tdata.input)))
		assert.Equal(t, tdata.excepted, largestValues2(tutils.IntSliceToBinaryTree(tdata.input)))
	}
}
