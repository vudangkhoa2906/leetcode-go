package palindromepartitioningii

import "vudangkhoa2906-leetcode-go/common"

func minCut(s string) int {
	lenS := len(s)
	dp := make([]int, lenS)
	var tempRes int

	for highIdx := 0; highIdx < lenS; highIdx++ {
		if common.IsPalindromicSubstring(s, 0, highIdx) {
			dp[highIdx] = 1
		} else {
			tempRes = highIdx + 1
			for lowIdx := 0; lowIdx < highIdx; lowIdx++ {
				if common.IsPalindromicSubstring(s, lowIdx+1, highIdx) {
					if tempRes > 1+dp[lowIdx] {
						tempRes = 1 + dp[lowIdx]
					}
				}
			}
			dp[highIdx] = tempRes
		}
	}
	return dp[lenS-1] - 1
}
