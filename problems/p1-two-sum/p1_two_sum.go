package two_sum

/*
1. 两数之和

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/two-sum/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		m[v] = i
	}

	for i, v := range nums {
		if index, ok := m[target-v]; ok && index != i {
			return []int{i, index}
		}
	}
	return nil
}
