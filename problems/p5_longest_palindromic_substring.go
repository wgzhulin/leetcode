package problems

/*
5. 最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。

示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

示例 2：
输入：s = "cbbd"
输出："bb"

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
