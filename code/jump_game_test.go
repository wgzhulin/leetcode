package code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
示例 1：
输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。

示例 2：
输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。

*/
func TestCanJump(t *testing.T) {
	type data struct {
		input  []int
		except bool
	}

	testData := []data{
		{
			input:  []int{2, 3, 1, 1, 4},
			except: true,
		},
		{
			input:  []int{3, 2, 1, 0, 4},
			except: false,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, canJump(tdata.input))
	}
}
