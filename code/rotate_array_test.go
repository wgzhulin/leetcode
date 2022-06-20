package code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
示例 1:
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]

示例 2:
输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]
*/
func TestRotate(t *testing.T) {
	type data struct {
		input  []int
		input2 int
		except []int
	}

	testData := []data{
		{
			input:  []int{1, 2, 3, 4, 5, 6, 7},
			input2: 3,
			except: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			input:  []int{-1, -100, 3, 99},
			input2: 2,
			except: []int{3, 99, -1, -100},
		},
	}

	for _, tdata := range testData {
		rotate(tdata.input, tdata.input2)
		assert.Equal(t, tdata.except, tdata.input)
	}
}
