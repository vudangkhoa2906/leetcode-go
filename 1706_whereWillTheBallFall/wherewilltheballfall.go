package wherewilltheballfall

import "vudangkhoa2906-leetcode-go/common"

func findBall(grid [][]int) []int {
	numRows, numCols := len(grid), len(grid[0])
	res := make([]int, numCols)

	var findBallRec func(pos common.Pos) common.Pos
	findBallRec = func(pos common.Pos) common.Pos {
		if pos.R == numRows-1 || pos.C == -1 {
			return pos
		} else if pos.C > 0 && grid[pos.R+1][pos.C-1] == -1 && grid[pos.R+1][pos.C] == -1 {
			return findBallRec(common.Pos{R: pos.R + 1, C: pos.C - 1})
		} else if pos.C < numCols-1 && grid[pos.R+1][pos.C+1] == 1 && grid[pos.R+1][pos.C] == 1 {
			return findBallRec(common.Pos{R: pos.R + 1, C: pos.C + 1})
		} else {
			return common.Pos{R: pos.R, C: -1}
		}
	}

	for idx := range res {
		res[idx] = findBallRec(common.Pos{R: -1, C: idx}).C
	}
	return res
}
