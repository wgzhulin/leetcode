package binary_search

/*
704. 二分查找

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/binary-search/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		midValue := nums[mid]
		if midValue == target {
			return mid
		} else if midValue < target {
			left = mid + 1
		} else {
			right = right - 1
		}
	}
	return -1
}
