package best_poker_hand

/*
2347. 最好的扑克手牌

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/best-poker-hand/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func bestHand(ranks []int, suits []byte) string {
	m1 := make(map[string]struct{}, len(suits))
	for _, v := range suits {
		m1[string(v)] = struct{}{}
	}

	m2 := make(map[int]int, len(ranks))
	for _, v := range ranks {
		m2[v] = m2[v] + 1
	}

	if len(m1) == 1 {
		return "Flush"
	}

	max := 0
	for _, v := range m2 {
		if v > max {
			max = v
		}
	}

	if max >= 3 {
		return "Three of a Kind"
	} else if max >= 2 {
		return "Pair"
	}

	return "High Card"
}
