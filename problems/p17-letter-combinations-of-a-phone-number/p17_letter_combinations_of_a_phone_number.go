package letter_combinations_of_a_phone_number

/*
17. 电话号码的字母组合

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func letterCombinations(digits string) []string {
	digitMap := map[uint8][]rune{
		'2': []rune{'a', 'b', 'c'},
		'3': []rune{'d', 'e', 'f'},
		'4': []rune{'g', 'h', 'i'},
		'5': []rune{'j', 'l', 'l'},
		'6': []rune{'m', 'n', 'o'},
		'7': []rune{'p', 'q', 'r', 's'},
		'8': []rune{'t', 'u', 'v'},
		'9': []rune{'w', 'x', 'y', 'z'},
	}

	input := make([][]rune, 0, 4)
	for i := range digits {
		input = append(input, digitMap[digits[i]])
	}
	ans := []string{}
	if len(input) == 0 {
		return ans
	}
	var dfs = func(input [][]rune, index int, comb []rune) {}

	dfs = func(input [][]rune, index int, comb []rune) {
		if index == len(input) {
			ans = append(ans, string(comb))
		} else {
			runes := digitMap[digits[index]]
			runesCount := len(runes)

			index++
			for i := 0; i < runesCount; i++ {
				dfs(input, index, append(comb, runes[i]))
			}
		}
	}

	dfs(input, 0, []rune{})

	return ans
}
