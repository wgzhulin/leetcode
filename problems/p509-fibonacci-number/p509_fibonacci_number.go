package fibonacci_number

/*
509. 斐波那契数

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/fibonacci-number/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fibDp(n int) int {
	dp := make([]int, n+1)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

