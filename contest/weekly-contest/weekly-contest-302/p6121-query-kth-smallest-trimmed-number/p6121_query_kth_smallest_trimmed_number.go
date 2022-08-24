package query_kth_smallest_trimmed_number

import (
	"sort"
)

/*
6121. 裁剪数字后查询第 K 小的数字

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/query-kth-smallest-trimmed-number/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
	ans := make([]int, 0, len(queries))
	for _, query := range queries {
		trimmed := make([]string, len(nums))
		for i := range nums {
			trimmed[i] = nums[i][len(nums[i])-query[1] : len(nums[i])]
		}
		ans = append(ans, smallestNumberIndex(trimmed, query[0]))
	}
	return ans
}

// 第N个最小值的索引
func smallestNumberIndex(trimmed []string, smallest int) int {
	type Key struct {
		str string
		i   int
	}

	s := make([]Key, len(trimmed))
	for i := range trimmed {
		s[i] = Key{
			str: trimmed[i],
			i:   i,
		}
	}

	sort.SliceStable(s, func(i, j int) bool {
		return s[i].str < s[j].str
	})

	return s[smallest-1].i
}
