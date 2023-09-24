package editdistance

import "vudangkhoa2906-leetcode-go/common"

func minDistance(word1 string, word2 string) int {
	lenWord1, lenWord2 := len(word1), len(word2)
	if lenWord1 == 0 || lenWord2 == 0 {
		return lenWord1 + lenWord2
	}

	dp := make([][]int, lenWord1)
	for idx := range dp {
		dp[idx] = make([]int, lenWord2)
	}
	for idx1 := 0; idx1 < lenWord1; idx1++ {
		for idx2 := 0; idx2 < lenWord2; idx2++ {
			if word1[idx1] == word2[idx2] {
				if idx1 > 0 && idx2 > 0 {
					dp[idx1][idx2] = dp[idx1-1][idx2-1]
				} else {
					dp[idx1][idx2] = idx1 + idx2
				}
			} else {
				if idx1 > 0 && idx2 > 0 {
					dp[idx1][idx2] = 1 + common.Min(dp[idx1-1][idx2-1], common.Min(dp[idx1-1][idx2], dp[idx1][idx2-1]))
				} else if idx1 > 0 {
					dp[idx1][idx2] = 1 + dp[idx1-1][idx2]
				} else if idx2 > 0 {
					dp[idx1][idx2] = 1 + dp[idx1][idx2-1]
				} else {
					dp[idx1][idx2] = 1
				}
			}
		}
	}
	return dp[lenWord1-1][lenWord2-1]
}
