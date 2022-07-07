package problems

import (
	"github.com/zhulinw/leetcode/testify/assert"
	"testing"
)

/*
示例 1：
输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶

示例 2：
输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶
*/
func TestClimbStairs(t *testing.T) {
	type data struct {
		input  int
		except int
	}

	testData := []data{
		{
			input:  2,
			except: 2,
		},
		{
			input:  3,
			except: 3,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, climbStairs(tdata.input))
		assert.Equal(t, tdata.except, climbStairs2(tdata.input))
		assert.Equal(t, tdata.except, climbStairs3(tdata.input))
	}
}
