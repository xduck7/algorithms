package ctwobackpacks

func TwoBackpacks(n, w1, w2 int64, items [][2]int64) int64 {
	dp := make([][]int64, w1+1)
	for i := range dp {
		dp[i] = make([]int64, w2+1)
	}

	for k := 0; k < int(n); k++ {
		wi := items[k][0]
		vi := items[k][1]

		for i := w1; i >= 0; i-- {
			for j := w2; j >= 0; j-- {
				if i >= wi {
					if dp[i][j] < dp[i-wi][j]+vi {
						dp[i][j] = dp[i-wi][j] + vi
					}
				}
				if j >= wi {
					if dp[i][j] < dp[i][j-wi]+vi {
						dp[i][j] = dp[i][j-wi] + vi
					}
				}
			}
		}
	}

	return dp[w1][w2]
}
