package countsortedvowelstrings

func countVowelStrings(n int) int {
	dp := make([][]int, n)
	for idx := range dp {
		dp[idx] = make([]int, 5)
	}
	for idx := 0; idx < 5; idx++ {
		dp[0][idx] = 1
	}
	for idx1 := 1; idx1 < n; idx1++ {
		for idx2 := 0; idx2 < 5; idx2++ {
			for idx3 := 0; idx3 <= idx2; idx3++ {
				dp[idx1][idx2] += dp[idx1-1][idx3]
			}
		}
	}
	var res int
	for idx := 0; idx < 5; idx++ {
		res += dp[n-1][idx]
	}
	return res
}
