package problems

/*
26. 删除有序数组中的重复项
给你一个 升序排列的数组nums ，请你原地删除重复出现的元素，使每个元素只出现一次 ，返回删除后数组的新长度。元素的相对顺序应该保持一致 。
由于在某些语言中不能改变数组的长度，所以必须将结果放在数组nums的第一部分。更规范地说，如果在删除重复项之后有 k 个元素，那么nums的前k个元素应该保存最终结果。
将最终结果插入的前k个位置后返回k。
不要使用额外的空间，你必须在原地修改输入数组并在使用O(1)额外空间的条件下完成。

示例 1：
输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。

示例 2：
输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。

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
