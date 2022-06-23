package problems

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
示例 1：
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2

示例 2：
输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
*/
func TestFindMedianSortedArrays(t *testing.T) {
	type data struct {
		input1 []int
		input2 []int
		except float64
	}

	testData := []data{
		{
			input1: []int{1, 3},
			input2: []int{2},
			except: 2.00000,
		},
		{
			input1: []int{1, 2},
			input2: []int{3, 4},
			except: 2.50000,
		},
		{
			input1: []int{},
			input2: []int{1, 2, 3, 4, 5},
			except: 3,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, findMedianSortedArrays(tdata.input1, tdata.input2))
	}
}
