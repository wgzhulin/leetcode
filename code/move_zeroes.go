package code

/*
283. 移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例 1:
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]

示例 2:
输入: nums = [0]
输出: [0]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/move-zeroes/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func moveZeroes(nums []int) {
	temp := nums[:0]
	for i := range nums {
		if nums[i] != 0 {
			temp = append(temp, nums[i])
		}
	}

	temp = append(temp, make([]int, len(nums)-len(temp))...)
	copy(nums, temp)
}

func moveZeroes2(nums []int) {
	left, right, length := 0, 0, len(nums)

	for right < length {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
