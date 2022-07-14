package problems

/*
66. 加一

给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1：
输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。

示例 2：
输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。

示例 3：
输入：digits = [0]
输出：[1]

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
