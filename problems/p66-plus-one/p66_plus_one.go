package plus_one

/*
66. 加一

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/plus-one/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func plusOne(digits []int) []int {
	// 非空数组
	end := digits[len(digits)-1]
	carry := (end + 1) / 10
	digits[len(digits)-1] = (end + 1) % 10
	for i := len(digits) - 2; i >= 0; i-- {
		v := digits[i]
		digits[i] = (v + carry) % 10
		carry = (v + carry) / 10
	}
	if carry > 0 {
		s := make([]int, len(digits)+1)
		copy(s[1:], digits)
		s[0] = carry
		return s
	}
	return digits
}

func plusOne2(digits []int) []int {
	// 找到数组中第一个不为9的
	for i := len(digits) - 1; i >= 0; i-- {
		if (digits[i]+1)%10 != 0 {
			digits[i] = digits[i] + 1
			return digits
		} else {
			digits[i] = 0
		}
	}

	// digits中全为9
	s := make([]int, len(digits)+1)
	s[0] = 1
	return s
}
