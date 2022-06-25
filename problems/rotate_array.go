package problems

/*
189. 轮转数组
给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

示例 1:
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]

示例 2:
输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/rotate-array
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