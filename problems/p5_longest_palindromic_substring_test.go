package problems

import (
	"github.com/zhulinw/leetcode/testify/assert"
	"testing"
)

/*
示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

示例 2：
输入：s = "cbbd"
输出："bb"
*/
func TestLongestPalindrome(t *testing.T) {
	type data struct {
		input    string
		excepted string
	}

	testData := []data{
		{
			input:    "babad",
			excepted: "bab",
		},
		{
			input:    "cbbd",
			excepted: "bb",
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.excepted, longestPalindrome(tdata.input))
	}
}
