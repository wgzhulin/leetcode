package valid_parentheses

/*
20. 有效的括号

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/valid-parentheses/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func isValid(s string) bool {
	q := make([]byte, 0)

	q = append(q, s[0])
	for i := 1; i < len(s); i++ {
		// 匹配上队列中最后一个元素
		if len(q) > 0 && isParentThese(q[len(q)-1], s[i]) {
			q = q[:len(q)-1]
		} else {
			q = append(q, s[i])
		}
	}
	return len(q) == 0
}

// use hash table maybe fast?
func isParentThese(a, b byte) bool {
	if a == '[' && b == ']' {
		return true
	}
	if a == '(' && b == ')' {
		return true
	}
	if a == '{' && b == '}' {
		return true
	}

	return false
}