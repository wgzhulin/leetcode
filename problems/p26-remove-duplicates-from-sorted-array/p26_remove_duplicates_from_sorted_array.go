package remove_duplicates_from_sorted_array

/*
26. 删除有序数组中的重复项

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func removeDuplicates(nums []int) int {
	// O(n*n)?
	for i := 1; i < len(nums); {
		if nums[i] == nums[i-1] {
			for j := i + 1; j < len(nums); j++ {
				nums[j-1] = nums[j]
			}
			nums = nums[:len(nums)-1]
		} else {
			i++
		}
	}
	return len(nums)
}

// 双指针
func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	fast, slow := 1, 1
	for fast < len(nums) {
		if nums[fast] != nums[slow-1] {
			nums[slow] = nums[fast]
			slow++
		}

		fast++
	}
	return len(nums[:slow])
}

