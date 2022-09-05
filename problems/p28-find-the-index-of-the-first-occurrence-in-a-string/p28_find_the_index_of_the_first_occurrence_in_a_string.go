package find_the_index_of_the_first_occurrence_in_a_string

/*
28. 实现 strStr()

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// 滑动窗口
func strStr(haystack string, needle string) int {
	step := len(needle)

	for j := 0; j+step <= len(haystack); j++ {
		t := haystack[j : j+step]
		if t == needle {
			return j
		}
	}

	return -1
}

func strStr2(haystack string, needle string) int {
	for i := range haystack {
		if haystack[i] == needle[0] {
			if len(needle) <= len(haystack[i:]) && haystack[i : i+len(needle)] == needle {
				return i
			}
		}
	}

	return -1
}
