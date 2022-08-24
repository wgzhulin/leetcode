package make_array_zero_by_subtracting_equal_amounts

/*
2357. 使数组中所有元素都等于零

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/make-array-zero-by-subtracting-equal-amounts/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func minimumOperations(nums []int) int {
	m := make(map[int]struct{}, len(nums))

	for i := range nums {
		m[nums[i]] = struct{}{}
	}
	delete(m, 0)
	return len(m)
}
