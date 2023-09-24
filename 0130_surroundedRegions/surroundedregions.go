package surroundedregions

import "vudangkhoa2906-leetcode-go/common"

func solve(board [][]byte) {
	numRows, numCols := len(board), len(board[0])
	if numRows <= 2 || numCols <= 2 {
		return
	}
	marked := make([][]bool, numRows)
	for idx := range marked {
		marked[idx] = make([]bool, numCols)
	}
	getInitialPeripherl := func() []common.Pos {
		var res []common.Pos
		for c := 0; c < numCols; c++ {
			if board[0][c] == 'O' {
				res = append(res, common.Pos{R: 0, C: c})
			}
			if board[numRows-1][c] == 'O' {
				res = append(res, common.Pos{R: numRows - 1, C: c})
			}
		}
		for r := 1; r < numRows-1; r++ {
			if board[r][0] == 'O' {
				res = append(res, common.Pos{R: r, C: 0})
			}
			if board[r][numCols-1] == 'O' {
				res = append(res, common.Pos{R: r, C: numCols - 1})
			}
		}
		return res
	}

	var depthFirstSearch func(pos common.Pos)
	depthFirstSearch = func(pos common.Pos) {
		marked[pos.R][pos.C] = true
		var adjRow, adjCol int
		delta := [][]int{{-1, 1, 0, 0}, {0, 0, -1, 1}}
		for idx := 0; idx < 4; idx++ {
			adjRow, adjCol = pos.R+delta[0][idx], pos.C+delta[1][idx]
			if adjRow >= 0 && adjRow < numRows && adjCol >= 0 && adjCol < numCols && board[adjRow][adjCol] == 'O' && !marked[adjRow][adjCol] {
				depthFirstSearch(common.Pos{R: adjRow, C: adjCol})
			}
		}
	}

	for _, pos := range getInitialPeripherl() {
		depthFirstSearch(pos)
	}

	for r := 1; r < numRows-1; r++ {
		for c := 1; c < numCols-1; c++ {
			if board[r][c] == 'O' && !marked[r][c] {
				board[r][c] = 'X'
			}
		}
	}
}
