package cherrypickup

import "vudangkhoa2906-leetcode-go/common"

func cherryPickup(grid [][]int) int {
	lenGrid := len(grid)
	dp := make([][][]int, lenGrid)
	for idx1 := range dp {
		dp[idx1] = make([][]int, lenGrid)
		for idx2 := range dp[idx1] {
			dp[idx1][idx2] = make([]int, lenGrid)
		}
	}
	var c2, cur, prev int
	for r1 := 0; r1 < lenGrid; r1++ {
		for c1 := 0; c1 < lenGrid; c1++ {
			for r2 := 0; r2 < lenGrid; r2++ {
				if r1 == 0 && c1 == 0 && r2 == 0 {
					dp[r1][c1][r2] = grid[0][0]
					continue
				}
				c2 = r1 + c1 - r2
				if c2 < 0 || c2 >= lenGrid {
					continue
				}
				if grid[r1][c1] == -1 || grid[r2][c2] == -1 {
					dp[r1][c1][r2] = -1
					continue
				}
				prev = -1
				if r1 > 0 && r2 > 0 && dp[r1-1][c1][r2-1] > prev {
					prev = dp[r1-1][c1][r2-1]
				}
				if r1 > 0 && c2 > 0 && dp[r1-1][c1][r2] > prev {
					prev = dp[r1-1][c1][r2]
				}
				if c1 > 0 && r2 > 0 && dp[r1][c1-1][r2-1] > prev {
					prev = dp[r1][c1-1][r2-1]
				}
				if c1 > 0 && c2 > 0 && dp[r1][c1-1][r2] > prev {
					prev = dp[r1][c1-1][r2]
				}
				if prev == -1 {
					dp[r1][c1][r2] = -1
					continue
				}
				if r1 == r2 {
					cur = grid[r1][c1]
				} else {
					cur = grid[r1][c1] + grid[r2][c2]
				}
				dp[r1][c1][r2] = cur + prev
			}
		}
	}
	return common.Max(dp[lenGrid-1][lenGrid-1][lenGrid-1], 0)
}
