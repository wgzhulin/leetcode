package best_time_to_buy_and_sell_stock

/*
121. 买卖股票的最佳时机

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func maxProfit(prices []int) int {
	temp, ans := 0, 0

	for i := 1; i < len(prices); i++ {
		temp = temp + prices[i] - prices[i-1]
		temp = max(0, temp)  // 短周期内最大利润
		ans = max(ans, temp) // 记录整个周期内最大利润
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit2(prices []int) int {
	ans := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if p := prices[j] - prices[i]; p > ans {
				ans = p
			}
		}
	}

	return ans
}
