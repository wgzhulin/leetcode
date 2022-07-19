package longest_palindrome

/*
409. 最长回文串

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-palindrome/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func longestPalindrome(s string) int {
	m := make(map[uint8]int, len(s))
	for i := range s {
		m[s[i]] = m[s[i]] + 1
	}

	count := 0
	for _, v := range m {
		count = count + v/2
	}

	ans := count * 2
	if ans < len(s) {
		return ans + 1
	}

	return ans
}
