package problems

import (
	"github.com/zhulinw/leetcode/testify/assert"
	"testing"
)

/*
示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/
func TestLengthOfLongestSubstring(t *testing.T) {
	type data struct {
		input  string
		except int
	}

	testData := []data{
		{
			input:  "abcabcbb",
			except: 3,
		},
		{
			input:  "bbbbb",
			except: 1,
		},
		{
			input:  "pwwkew",
			except: 3,
		},
		{
			input:  " ",
			except: 1,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, lengthOfLongestSubstring(tdata.input))
	}
}
