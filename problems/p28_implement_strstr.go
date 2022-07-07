package problems

/*
28. 实现 strStr()
实现 strStr() 函数。

给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。

示例 1：
输入：haystack = "hello", needle = "ll"
输出：2

示例 2：
输入：haystack = "aaaaa", needle = "bba"
输出：-1


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/implement-strstr/
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
