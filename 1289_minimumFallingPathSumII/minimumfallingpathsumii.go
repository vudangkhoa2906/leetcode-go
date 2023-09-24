package minimumfallingpathsumii

import "math"

func minFallingPathSum(grid [][]int) int {
	res := math.MaxInt
	numRows, numCols := len(grid), len(grid[0])
	resCur, resPrev := make([]int, numCols), make([]int, numCols)
	copy(resPrev, grid[0])
	var tempPrev int
	for r := 1; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			tempPrev = math.MaxInt
			for c1 := 0; c1 < numCols; c1++ {
				if c != c1 && resPrev[c1] < tempPrev {
					tempPrev = resPrev[c1]
				}
			}
			resCur[c] = tempPrev + grid[r][c]
		}
		copy(resPrev, resCur)
	}
	for c := 0; c < numCols; c++ {
		if resPrev[c] < res {
			res = resPrev[c]
		}
	}
	return res
}
