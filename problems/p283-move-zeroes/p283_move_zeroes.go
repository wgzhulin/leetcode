package move_zeroes

/*
283. 移动零

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
