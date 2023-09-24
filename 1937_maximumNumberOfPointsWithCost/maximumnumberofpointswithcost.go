package maximumnumberofpointswithcost

import (
	"math"
	"vudangkhoa2906-leetcode-go/common"
)

func maxPoints(points [][]int) int64 {
	res := int64(math.MinInt64)
	numRows, numCols := len(points), len(points[0])
	resCur, resPrev := make([]int64, numCols), make([]int64, numCols)
	for c := 0; c < numCols; c++ {
		resPrev[c] = int64(points[0][c])
	}
	var tempPrev int64
	for r := 1; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			tempPrev = math.MinInt64
			for c1 := 0; c1 < numCols; c1++ {
				if resPrev[c1]-int64(common.Abs(c-c1)) > tempPrev {
					tempPrev = resPrev[c1] - int64(common.Abs(c-c1))
				}
			}
			resCur[c] = int64(points[r][c]) + tempPrev
		}
		copy(resPrev, resCur)
	}
	for c := 0; c < numCols; c++ {
		if resPrev[c] > res {
			res = resPrev[c]
		}
	}
	return res
}
