package length_of_last_word

/*
58. 最后一个单词的长度

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/length-of-last-word/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func lengthOfLastWord(s string) int {
	start, end := 0, len(s)-1
	endFlag := true // 记录设置了end
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' && endFlag {
			end = i
			endFlag = false
		}

		if i-1 >= 0 && s[i] != ' ' && s[i-1] == ' ' && !endFlag {
			start = i
			break
		}
	}
	return len(s[start : end+1])
}

func lengthOfLastWord2(s string) int {
	endIndex := len(s) - 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			endIndex = i
			break
		}
	}

	startIndex := 0
	for i := endIndex; i >= 0; i-- {
		if s[i] == ' ' {
			startIndex = i + 1
			break
		}
	}
	return len(s[startIndex:endIndex+1])
}

