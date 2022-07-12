package problems

/*
35. 搜索插入位置
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。请必须使用时间复杂度为O(logn)的算法。

示例 1:
输入: nums = [1,3,5,6], target = 5
输出: 2

示例 2:
输入: nums = [1,3,5,6], target = 2
输出: 1

示例 3:
输入: nums = [1,3,5,6], target = 7
输出: 4

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/search-insert-position/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func searchInsert(nums []int, target int) int {
	ans := 0
	for i := range nums {
		if nums[i] >= target {
			ans = i
			break
		} else if nums[i] < target {
			ans++
		}
	}
	return ans
}

func searchInsert2(nums []int, target int) int {
	// 有序数组 二分查找
	left, right := 0, len(nums)-1
	ans := 0
	for left <= right {
		mid := left + (right-left)/2
		ans = mid
		if nums[mid] == target {
			break
		} else if nums[mid] < target {
			left = mid + 1
			// 比有序数据最大的元素还大，需要加一
			ans = ans + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
