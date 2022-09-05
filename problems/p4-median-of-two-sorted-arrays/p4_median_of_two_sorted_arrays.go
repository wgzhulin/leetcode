package median_of_two_sorted_arrays

import "sort"

/*
4. 寻找两个正序数组的中位数

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/median-of-two-sorted-arrays/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	s := append(nums1, nums2...)
	if len(s) == 1 {
		return float64(s[0])
	}
	if len(s) == 2 {
		return float64(s[0]+s[1]) / 2
	}
	sort.Ints(s)

	l := len(s) / 2
	if len(s)%2 == 0 {
		return float64(s[l-1]+s[l]) / 2
	}
	return float64(s[l])
}