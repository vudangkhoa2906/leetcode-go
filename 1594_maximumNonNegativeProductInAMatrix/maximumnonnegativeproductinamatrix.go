package maximumnonnegativeproductinamatrix

import "vudangkhoa2906-leetcode-go/common"

func maxProductPath(grid [][]int) int {
	numRows, numCols := len(grid), len(grid[0])
	dp1, dp2 := make([][]int, numRows), make([][]int, numRows)
	for idx := 0; idx < numRows; idx++ {
		dp1[idx], dp2[idx] = make([]int, numCols), make([]int, numCols)
	}
	var prevMax, prevMin int
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			if r == 0 && c == 0 {
				dp1[r][c], dp2[r][c] = grid[r][c], grid[r][c]
				continue
			}
			if r > 0 && c > 0 {
				prevMax = common.Max(dp1[r-1][c], dp1[r][c-1])
				prevMin = common.Min(dp2[r-1][c], dp2[r][c-1])
			} else if r == 0 {
				prevMax = dp1[r][c-1]
				prevMin = dp2[r][c-1]
			} else {
				prevMax = dp1[r-1][c]
				prevMin = dp2[r-1][c]
			}
			dp1[r][c] = common.Max(prevMax*grid[r][c], prevMin*grid[r][c])
			dp2[r][c] = common.Min(prevMax*grid[r][c], prevMin*grid[r][c])
		}
	}
	return common.Max(dp1[numRows-1][numCols-1], -1)
}
