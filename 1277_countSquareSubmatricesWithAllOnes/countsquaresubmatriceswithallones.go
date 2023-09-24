package countsquaresubmatriceswithallones

import (
	"vudangkhoa2906-leetcode-go/common"
)

func countSquares(matrix [][]int) int {
	numRows, numCols := len(matrix), len(matrix[0])
	dp := make([][]int, numRows)
	var res int
	for idx := range dp {
		dp[idx] = make([]int, numCols)
	}
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			if matrix[r][c] == 1 {
				if r > 0 && c > 0 {
					dp[r][c] = 1 + common.Min(dp[r-1][c-1], common.Min(dp[r-1][c], dp[r][c-1]))
				} else {
					dp[r][c] = 1
				}
			}
			res += dp[r][c]
		}
	}
	return res
}
