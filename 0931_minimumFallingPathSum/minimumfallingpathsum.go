package minimumfallingpathsum

import "math"

func minFallingPathSum(matrix [][]int) int {
	res := math.MaxInt
	numRows, numCols := len(matrix), len(matrix[0])
	resCur, resPrev := make([]int, numCols), make([]int, numCols)
	copy(resPrev, matrix[0])
	deltaC := []int{-1, 0, 1}
	var tempPrev int
	for r := 1; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			tempPrev = math.MaxInt
			for i := 0; i < 3; i++ {
				if c+deltaC[i] >= 0 && c+deltaC[i] < numCols && resPrev[c+deltaC[i]] < tempPrev {
					tempPrev = resPrev[c+deltaC[i]]
				}
			}
			resCur[c] = matrix[r][c] + tempPrev
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
