package problems

import (
	"github.com/zhulinw/leetcode/testify/assert"
	"github.com/zhulinw/leetcode/testify/tutils"
	"testing"
)

/*
示例 1：
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

示例 2：
输入：l1 = [0], l2 = [0]
输出：[0]

示例 3：
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]
*/
func TestAddTwoNumbers(t *testing.T) {
	type data struct {
		input1 []int
		input2 []int
		except []int
	}

	testData := []data{
		{
			input1: []int{2, 4, 3},
			input2: []int{5, 6, 4},
			except: []int{7, 0, 8},
		},
		{
			input1: []int{0},
			input2: []int{0},
			except: []int{0},
		},
		{
			input1: []int{9, 9, 9, 9, 9, 9, 9},
			input2: []int{9, 9, 9, 9},
			except: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tutils.SliceToListNode(tdata.except), addTwoNumbers(tutils.SliceToListNode(tdata.input1), tutils.SliceToListNode(tdata.input2)))
	}
}
