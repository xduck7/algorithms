package countsplitting

func CountSplitting(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for part := 1; part <= n; part++ {
		for i := part; i <= n; i++ {
			dp[i] += dp[i-part]
		}
	}
	return dp[n]
}
