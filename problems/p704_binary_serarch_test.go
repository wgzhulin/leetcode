package problems

import (
	"github.com/zhulinw/leetcode/testify/assert"
	"testing"
)

/*
示例 1:
输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4

示例 2:
输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1
*/
func TestSearch(t *testing.T) {
	type data struct {
		input1 []int
		input2 int
		except int
	}

	testData := []data{
		{
			input1: []int{-1, 0, 3, 5, 9, 12},
			input2: 9,
			except: 4,
		},
		{
			input1: []int{-1, 0, 3, 5, 9, 12},
			input2: 2,
			except: -1,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, search(tdata.input1, tdata.input2))
	}
}
