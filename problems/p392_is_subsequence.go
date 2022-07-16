package problems

/*
392. 判断子序列

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/is-subsequence/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func isSubsequence(s string, t string) bool {
	index := 0
	for i := range t {
		if index >= len(s) {
			return true
		}
		// 可以优化一下，已经找到子序列直接返回
		if index < len(s) && t[i] == s[index] {
			index++
		}
	}

	return len(s) == index
}

func isSubsequence2(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	return len(s) == i
}
