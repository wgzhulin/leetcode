package problems

import (
	"github.com/zhulinw/leetcode/testify/assert"
	"testing"
)

/*
示例 1：
输入：strs = ["flower","flow","flight"]
输出："fl"

示例 2：
输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。
*/
func TestLongestCommonPrefix(t *testing.T) {
	type data struct {
		input  []string
		except string
	}

	testData := []data{
		{
			input:  []string{"flower", "flow", "flight"},
			except: "fl",
		},
		{
			input:  []string{"dog", "racecar", "car"},
			except: "",
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, longestCommonPrefix(tdata.input))
	}
}
