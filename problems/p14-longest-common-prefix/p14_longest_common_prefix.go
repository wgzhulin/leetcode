package longest_common_prefix

import "strings"

/*
14. 最长公共前缀

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-common-prefix/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	prefix := strs[0]
	for i := 1; i <= len(prefix); i++ {
		for _, str := range strs {
			if !strings.HasPrefix(str, prefix[:i]) {
				return prefix[:i-1]
			}
		}
	}
	return prefix
}
