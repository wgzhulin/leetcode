package rotate_array

/*
189. 轮转数组

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/rotate-array/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func rotate(nums []int, k int) {
	k = k % len(nums)
	tmp := make([]int, len(nums)-k)
	copy(tmp, nums[:len(nums)-k])

	nums = append(nums[0:0], nums[len(nums)-k:]...)
	nums = append(nums, tmp...)
}

func rotate2(nums []int, k int) {
	k = k % len(nums)

	length := len(nums)
	tmp := make([]int, length)
	for i, v := range nums {
		tmp[(i+k)%length] = v
	}
	copy(nums, tmp)
}
