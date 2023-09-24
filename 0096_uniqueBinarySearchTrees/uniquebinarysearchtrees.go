package uniquebinarysearchtrees

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	var numTreesRec func(count int) int
	numTreesRec = func(count int) int {
		if dp[count] != 0 {
			return dp[count]
		}
		var res int
		for leftCount := 0; leftCount <= count-1; leftCount++ {
			res += numTreesRec(leftCount) * numTreesRec(count-1-leftCount)
		}
		dp[count] = res
		return res
	}

	return numTreesRec(n)
}
