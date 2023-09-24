package detonatethemaximumbombs

import "vudangkhoa2906-leetcode-go/common"

func maximumDetonation(bombs [][]int) int {
	lenBombs := len(bombs)
	marked := make([]bool, lenBombs)
	var passedBy []bool
	var tempRes, res int

	var depthFirstSearch func(idx int)
	depthFirstSearch = func(idx int) {
		if passedBy[idx] {
			return
		}
		tempRes++
		marked[idx], passedBy[idx] = true, true
		for idxOther := range bombs {
			if passedBy[idxOther] {
				continue
			}
			if common.Square(bombs[idx][0]-bombs[idxOther][0])+common.Square(bombs[idx][1]-bombs[idxOther][1]) <= common.Square(bombs[idx][2]) {
				depthFirstSearch(idxOther)
			}
		}
	}

	for idx := range bombs {
		if !marked[idx] {
			passedBy = make([]bool, lenBombs)
			marked[idx] = true
			tempRes = 0
			depthFirstSearch(idx)
			if tempRes > res {
				res = tempRes
			}
		}
	}
	return res
}
