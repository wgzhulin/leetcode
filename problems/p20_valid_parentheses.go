package problems

/*
20. 有效的括号

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
 
示例 1：
输入：s = "()"
输出：true

示例 2：
输入：s = "()[]{}"
输出：true

示例 3：
输入：s = "(]"
输出：false

示例 4：
输入：s = "([)]"
输出：false

示例 5：
输入：s = "{[]}"
输出：true

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/valid-parentheses
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
