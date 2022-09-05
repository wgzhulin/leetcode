package running_sum_of_1d_array

/*
1480. 一维数组的动态和

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/running-sum-of-1d-array/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}
	return nums
}

