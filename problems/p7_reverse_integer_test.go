package problems

import (
	"github.com/zhulinw/leetcode/testify/assert"
	"testing"
)

/*
示例 1：
输入：x = 123
输出：321

示例 2：
输入：x = -123
输出：-321

示例 3：
输入：x = 120
输出：21

示例 4：
输入：x = 0
输出：0
*/
func TestReverse(t *testing.T) {
	type data struct {
		input    int
		excepted int
	}

	testData := []data{
		{
			input:    123,
			excepted: 321,
		},
		{
			input:    -123,
			excepted: -321,
		},
		{
			input:    120,
			excepted: 21,
		},
		{
			input:    0,
			excepted: 0,
		},
		{
			input:    1534236469,
			excepted: 0,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.excepted, reverse(tdata.input))
	}
}
