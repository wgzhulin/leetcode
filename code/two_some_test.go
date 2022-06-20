package code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

示例 2：
输入：nums = [3,2,4], target = 6
输出：[1,2]

示例 3：
输入：nums = [3,3], target = 6
输出：[0,1]
*/
func TestTwoSum(t *testing.T) {
	type data struct {
		input  []int
		input2 int
		except []int
	}

	testData := []data{
		{
			input:  []int{2, 7, 11, 15},
			input2: 9,
			except: []int{0, 1},
		},
		{
			input:  []int{3, 2, 4},
			input2: 6,
			except: []int{1, 2},
		},
		{
			input:  []int{3, 3},
			input2: 6,
			except: []int{0, 1},
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, twoSum(tdata.input, tdata.input2))
	}
}
