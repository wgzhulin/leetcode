package isomorphic_strings

/*
205. 同构字符串

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/isomorphic-strings/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	f1, f2 := true, true
	for i := 0; i < len(s); i++ {
		// t中相同的字符是否映射唯一s中唯一的字符
		for j := i + 1; j < len(t); j++ {
			if t[j] == t[i] && s[j] != s[i] {
				f1 = false
				break
			}
		}
		// s中相同的字符是否映射唯一t中唯一的字符
		for j := i + 1; j < len(s); j++ {
			if s[j] == s[i] && t[j] != t[i] {
				f2 = false
				break
			}
		}
	}
	return f1 && f2
}

func isIsomorphic2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	ms := make(map[uint8]uint8, len(s))
	mt := make(map[uint8]uint8, len(t))
	for i := range s {
		s1 := s[i]
		t1 := t[i]
		tv, okv := ms[s1]
		sv, okt := mt[t1]
		if (okv || okt) && (sv != s1 || tv != t1) {
			return false
		} else {
			ms[s1] = t1
			mt[t1] = s1
		}
	}
	return true
}