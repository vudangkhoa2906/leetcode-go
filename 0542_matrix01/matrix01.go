package matrix01

import "vudangkhoa2906-leetcode-go/common"

func updateMatrix(mat [][]int) [][]int {
	numRows := len(mat)
	numCols := len(mat[0])
	distance := 2
	var queue []common.Pos
	queue = append(queue, getInitialZeroes(mat, numRows, numCols)...)
	for levelCount := len(queue); levelCount > 0; levelCount = len(queue) {
		for _, pos := range queue[:levelCount] {
			if pos.R > 0 && mat[pos.R-1][pos.C] == 1 {
				mat[pos.R-1][pos.C] = distance
				queue = append(queue, common.Pos{R: pos.R - 1, C: pos.C})
			}
			if pos.C > 0 && mat[pos.R][pos.C-1] == 1 {
				mat[pos.R][pos.C-1] = distance
				queue = append(queue, common.Pos{R: pos.R, C: pos.C - 1})
			}
			if pos.R < numRows-1 && mat[pos.R+1][pos.C] == 1 {
				mat[pos.R+1][pos.C] = distance
				queue = append(queue, common.Pos{R: pos.R + 1, C: pos.C})
			}
			if pos.C < numCols-1 && mat[pos.R][pos.C+1] == 1 {
				mat[pos.R][pos.C+1] = distance
				queue = append(queue, common.Pos{R: pos.R, C: pos.C + 1})
			}
		}
		distance++
		queue = queue[levelCount:]
	}
	normalize(mat, numRows, numCols)
	return mat
}

func getInitialZeroes(mat [][]int, numRows, numCols int) []common.Pos {
	var res []common.Pos
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			if mat[r][c] == 0 {
				res = append(res, common.Pos{R: r, C: c})
			}
		}
	}
	return res
}

func normalize(mat [][]int, numRows, numCols int) {
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			if mat[r][c] > 0 {
				mat[r][c]--
			}
		}
	}
}
