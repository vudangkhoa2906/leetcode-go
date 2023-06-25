package cherrypickupii

import "math"

func cherryPickup(grid [][]int) int {
	var res int
	numRows := len(grid)
	numCols := len(grid[0])
	resCur, resPrev := make([][]int, numCols), make([][]int, numCols)
	for c := 0; c < numCols; c++ {
		resCur[c] = make([]int, numCols)
		resPrev[c] = make([]int, numCols)
	}
	for c1 := 0; c1 < numCols; c1++ {
		for c2 := 0; c2 < numCols; c2++ {
			resPrev[c1][c2] = math.MinInt
		}
	}
	resPrev[0][numCols-1] = grid[0][0] + grid[0][numCols-1]
	deltaC1 := []int{-1, 0, 1, -1, 0, 1, -1, 0, 1}
	deltaC2 := []int{-1, -1, -1, 0, 0, 0, 1, 1, 1}
	var tempCur, tempPrev int
	for r := 1; r < numRows; r++ {
		for c1 := 0; c1 < numCols; c1++ {
			for c2 := 0; c2 < numCols; c2++ {
				tempPrev = math.MinInt
				if c1 == c2 {
					tempCur = grid[r][c1]
				} else {
					tempCur = grid[r][c1] + grid[r][c2]
				}
				for i := 0; i < 9; i++ {
					if c1+deltaC1[i] >= 0 && c1+deltaC1[i] < numCols && c2+deltaC2[i] >= 0 && c2+deltaC2[i] < numCols && resPrev[c1+deltaC1[i]][c2+deltaC2[i]] > tempPrev {
						tempPrev = resPrev[c1+deltaC1[i]][c2+deltaC2[i]]
					}
				}
				resCur[c1][c2] = tempCur + tempPrev
			}
		}
		for c := 0; c < numCols; c++ {
			copy(resPrev[c], resCur[c])
		}
	}
	for c1 := 0; c1 < numCols; c1++ {
		for c2 := 0; c2 < numCols; c2++ {
			if resPrev[c1][c2] > res {
				res = resPrev[c1][c2]
			}
		}
	}
	return res
}
