package surroundedregions

import "vudangkhoa2906-leetcode-go/common"

func solve(board [][]byte) {
	numRows := len(board)
	numCols := len(board[0])
	marked := common.MarkedArray(numRows, numCols)
	for _, pos := range getInitialPeripheral(board, numRows, numCols) {
		depthFirstSearch(board, marked, numRows, numCols, pos)
	}
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			if board[r][c] == 'O' && !marked[r][c] {
				board[r][c] = 'X'
			}
		}
	}
}

func depthFirstSearch(board [][]byte, marked [][]bool, numRows, numCols int, s common.Pos) {
	marked[s.R][s.C] = true
	if s.R > 0 && !marked[s.R-1][s.C] && board[s.R-1][s.C] == 'O' {
		depthFirstSearch(board, marked, numRows, numCols, common.Pos{R: s.R - 1, C: s.C})
	}
	if s.R < numRows-1 && !marked[s.R+1][s.C] && board[s.R+1][s.C] == 'O' {
		depthFirstSearch(board, marked, numRows, numCols, common.Pos{R: s.R + 1, C: s.C})
	}
	if s.C > 0 && !marked[s.R][s.C-1] && board[s.R][s.C-1] == 'O' {
		depthFirstSearch(board, marked, numRows, numCols, common.Pos{R: s.R, C: s.C - 1})
	}
	if s.C < numCols-1 && !marked[s.R][s.C+1] && board[s.R][s.C+1] == 'O' {
		depthFirstSearch(board, marked, numRows, numCols, common.Pos{R: s.R, C: s.C + 1})
	}
}

func getInitialPeripheral(board [][]byte, numRows, numCols int) []common.Pos {
	var res []common.Pos
	for c := 0; c < numCols; c++ {
		if board[0][c] == 'O' {
			res = append(res, common.Pos{R: 0, C: c})
		}
		if numRows > 1 {
			if board[numRows-1][c] == 'O' {
				res = append(res, common.Pos{R: numRows - 1, C: c})
			}
		}
	}
	for r := 1; r < numRows-1; r++ {
		if board[r][0] == 'O' {
			res = append(res, common.Pos{R: r, C: 0})
		}
		if numCols > 1 {
			if board[r][numCols-1] == 'O' {
				res = append(res, common.Pos{R: r, C: numCols - 1})
			}
		}
	}
	return res
}
