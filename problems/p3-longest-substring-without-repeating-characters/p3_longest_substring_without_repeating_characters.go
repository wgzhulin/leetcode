package longest_substring_without_repeating_characters

/*
3. 无重复字符的最长子串

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-substring-without-repeating-characters/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// 滑动窗口暴力破解，部分用例超时
func lengthOfLongestSubstring(s string) int {
	ans := 0
	step := 1
	for i := 0; i < len(s); i++ {
		for j := 0; j+step < len(s)+1; j++ {
			if isStrNoRepeat(s[j:j+step]) && len(s[j:j+step]) > ans {
				ans = len(s[j : j+step])
			}
		}
		step++
	}
	return ans
}

func isStrNoRepeat(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}