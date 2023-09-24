package minimumpathcostinagrid

import "math"

func minPathCost(grid [][]int, moveCost [][]int) int {
	numRows, numCols := len(grid), len(grid[0])
	dp := make([][]int, numRows)
	for r := 0; r < numRows; r++ {
		dp[r] = make([]int, numCols)
	}
	var prev, tempPrev int
	res := math.MaxInt
	copy(dp[0], grid[0])
	for r := 1; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			prev = math.MaxInt
			for c1 := 0; c1 < numCols; c1++ {
				if tempPrev = dp[r-1][c1] + moveCost[grid[r-1][c1]][c]; tempPrev < prev {
					prev = tempPrev
				}
			}
			dp[r][c] = grid[r][c] + prev
		}
	}
	for c := 0; c < numCols; c++ {
		if dp[numRows-1][c] < res {
			res = dp[numRows-1][c]
		}
	}
	return res
}
