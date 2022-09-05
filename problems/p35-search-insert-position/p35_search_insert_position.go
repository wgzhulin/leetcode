package search_insert_position

/*
35. 搜索插入位置

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

