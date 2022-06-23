package example

// 空间换时间
func HasRepeatedCharactersInString(s string) bool {
	m := make(map[byte]struct{}, len(s))
	for i := range s {
		if _, ok := m[s[i]]; ok {
			return true
		}
		m[s[i]] = struct{}{}
	}
	return false
}

func HasRepeatedCharactersInString3(s string) bool {
	m := make(map[byte]struct{}, len(s))
	for i := range s {
		m[s[i]] = struct{}{}
	}
	return len(m) != len(s)
}

// 时间换空间
func HasRepeatedCharactersInString2(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return true
			}
		}
	}
	return false
}
