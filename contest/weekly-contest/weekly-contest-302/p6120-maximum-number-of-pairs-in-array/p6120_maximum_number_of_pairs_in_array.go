package maximum_number_of_pairs_in_array

/*
6120. 数组能形成多少数对

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximum-number-of-pairs-in-array/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func numberOfPairs(nums []int) []int {
	m := make(map[int]struct{}, len(nums))
	ans := make([]int, 2)
	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			ans[0]++
			delete(m, nums[i])
		}else {
			m[nums[i]] = struct{}{}
		}
	}
	ans[1] = len(m)
	return ans
}
