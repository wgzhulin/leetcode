package max_sum_of_a_pair_with_equal_sum_of_digits

/*
6164. 数位和相等数对的最大和

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/max-sum-of-a-pair-with-equal-sum-of-digits/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func maximumSum(nums []int) int {
	m := make(map[int][]int, len(nums))
	ans := -1
	for i := range nums {
		sum := bitSum(nums[i])
		m[sum] = append(m[sum], nums[i])
	}

	for _, v := range m {
		if len(v) < 2 {
			continue
		}
		ans = findSumMax(v, ans)
	}
	return ans
}

func findSumMax(s []int, ans int) int {
	for i := range s {
		for j := i + 1; j < len(s); j++ {
			if s[i]+s[j] > ans {
				ans = s[i] + s[j]
			}
		}
	}
	return ans
}

func bitSum(n int) int {
	sum := 0
	for n != 0 {
		sum += n % 10
		n = n / 10
	}
	return sum
}
