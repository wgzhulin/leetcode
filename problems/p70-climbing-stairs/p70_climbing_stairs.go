package climbing_stairs

/*
70. 爬楼梯

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/climbing-stairs/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// 利用上次的求和计算的结果
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	n1 := 1
	n2 := 2
	for i := 3; i <= n; i++ {
		temp := n1 + n2
		n1 = n2
		n2 = temp
	}
	return n2
}

// 递归
func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs2(n-1) + climbStairs2(n-2)
}

// 动态规划
func climbStairs3(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}

