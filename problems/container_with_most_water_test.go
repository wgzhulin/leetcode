package problems

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
示例 1：
输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

示例 2：
输入：height = [1,1]
输出：1
*/
func TestMaxArea(t *testing.T) {
	type data struct {
		input  []int
		except int
	}

	testData := []data{
		{
			input:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			except: 49,
		},
		{
			input:  []int{1, 1},
			except: 1,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, maxArea(tdata.input))
	}
}
