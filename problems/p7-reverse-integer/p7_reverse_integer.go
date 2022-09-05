package reverse_integer

import "math"

/*
7. 整数反转

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/reverse-integer/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func reverse(x int) int {
	ans := 0
	for x != 0 {
		t := x % 10
		ans = ans*10 + t
		x = x / 10
	}

	if ans > math.MaxInt32 || ans < math.MinInt32 {
		return 0
	}

	return ans
}