package longest_palindromic_substring

/*
5. 最长回文子串

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-palindromic-substring/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func longestPalindrome(s string) string {
	step := len(s)
	for i := 0; i < len(s); i++ {
		// s[i:j]；前包后不包
		for j := 0; j+step < len(s)+1; j++ {
			if isPalindrome(s[j : j+step]) {
				return s[j : j+step]
			}
		}
		step--
	}
	return ""
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

