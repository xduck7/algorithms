package backpack

func Backpack(n, w int, wm, vm []int) int {
	dp := make([]int, w+1)

	for i := 0; i < n; i++ {
		for j := w; j >= wm[i]; j-- {
			if dp[j] < dp[j-wm[i]]+vm[i] {
				dp[j] = dp[j-wm[i]] + vm[i]
			}
		}
	}

	return dp[w]
}
