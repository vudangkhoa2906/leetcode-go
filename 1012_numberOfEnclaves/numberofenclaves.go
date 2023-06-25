package numberofenclaves

import "vudangkhoa2906-leetcode-go/common"

func numEnclaves(grid [][]int) int {
	var res int
	numRows := len(grid)
	numCols := len(grid[0])
	marked := common.MarkedArray(numRows, numCols)
	for _, pos := range getInitialPeripheral(grid, numRows, numCols) {
		depthFirstSearch(grid, marked, numRows, numCols, pos)
	}
	for r := 1; r < numRows-1; r++ {
		for c := 1; c < numCols-1; c++ {
			if !marked[r][c] && grid[r][c] == 1 {
				res++
			}
		}
	}
	return res
}

func depthFirstSearch(grid [][]int, marked [][]bool, numRows, numCols int, s common.Pos) {
	marked[s.R][s.C] = true
	if s.R > 0 && !marked[s.R-1][s.C] && grid[s.R-1][s.C] == 1 {
		depthFirstSearch(grid, marked, numRows, numCols, common.Pos{R: s.R - 1, C: s.C})
	}
	if s.R < numRows-1 && !marked[s.R+1][s.C] && grid[s.R+1][s.C] == 1 {
		depthFirstSearch(grid, marked, numRows, numCols, common.Pos{R: s.R + 1, C: s.C})
	}
	if s.C > 0 && !marked[s.R][s.C-1] && grid[s.R][s.C-1] == 1 {
		depthFirstSearch(grid, marked, numRows, numCols, common.Pos{R: s.R, C: s.C - 1})
	}
	if s.C < numCols-1 && !marked[s.R][s.C+1] && grid[s.R][s.C+1] == 1 {
		depthFirstSearch(grid, marked, numRows, numCols, common.Pos{R: s.R, C: s.C + 1})
	}
}

func getInitialPeripheral(grid [][]int, numRows, numCols int) []common.Pos {
	var res []common.Pos
	for c := 0; c < numCols; c++ {
		if grid[0][c] == 1 {
			res = append(res, common.Pos{R: 0, C: c})
		}
		if numRows > 1 {
			if grid[numRows-1][c] == 1 {
				res = append(res, common.Pos{R: numRows - 1, C: c})
			}
		}
	}
	for r := 1; r < numRows-1; r++ {
		if grid[r][0] == 1 {
			res = append(res, common.Pos{R: r, C: 0})
		}
		if numCols > 1 {
			if grid[r][numCols-1] == 1 {
				res = append(res, common.Pos{R: r, C: numCols - 1})
			}
		}
	}
	return res
}
