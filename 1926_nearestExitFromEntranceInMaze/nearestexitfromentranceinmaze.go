package nearestexitfromentranceinmaze

import "vudangkhoa2906-leetcode-go/common"

func nearestExit(maze [][]byte, entrance []int) int {
	var res, adjRow, adjCol int
	numRows, numCols := len(maze), len(maze[0])
	queue := []common.Pos{{R: entrance[0], C: entrance[1]}}
	maze[entrance[0]][entrance[1]] = '+'
	delta := [][]int{{-1, 1, 0, 0}, {0, 0, -1, 1}}
	for count := len(queue); count > 0; count = len(queue) {
		res++
		for c := 0; c < count; c++ {
			for idx := 0; idx < 4; idx++ {
				adjRow, adjCol = queue[c].R+delta[0][idx], queue[c].C+delta[1][idx]
				if adjRow >= 0 && adjRow < numRows && adjCol >= 0 && adjCol < numCols && maze[adjRow][adjCol] == '.' {
					if adjRow == 0 || adjRow == numRows-1 || adjCol == 0 || adjCol == numCols-1 {
						return res
					}
					maze[adjRow][adjCol] = '+'
					queue = append(queue, common.Pos{R: adjRow, C: adjCol})
				}
			}
		}
		queue = queue[count:]
	}
	return -1
}
