package maximalrectangle

import "vudangkhoa2906-leetcode-go/common"

func maximalRectangle(matrix [][]byte) int {
	numRows, numCols := len(matrix), len(matrix[0])
	var stack []common.SliceElem
	var res int
	nle, ple, heights := make([]int, numCols), make([]int, numCols), make([]int, numCols)
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			if matrix[r][c] == '1' {
				heights[c]++
			} else {
				heights[c] = 0
			}
		}
		/* previous lesser element */
		for c := 0; c < numCols; c++ {
			for count := len(stack); count > 0 && stack[count-1].Val >= heights[c]; count = len(stack) {
				stack = stack[:count-1]
			}
			if count := len(stack); count > 0 {
				ple[c] = stack[count-1].Idx
			} else {
				ple[c] = -1
			}
			stack = append(stack, common.SliceElem{Idx: c, Val: heights[c]})
		}
		stack = stack[:0]
		/* next lesser element */
		for c := numCols - 1; c >= 0; c-- {
			for count := len(stack); count > 0 && stack[count-1].Val >= heights[c]; count = len(stack) {
				stack = stack[:count-1]
			}
			if count := len(stack); count > 0 {
				nle[c] = stack[count-1].Idx
			} else {
				nle[c] = numCols
			}
			stack = append(stack, common.SliceElem{Idx: c, Val: heights[c]})
		}
		stack = stack[:0]
		for c := 0; c < numCols; c++ {
			if tempRes := heights[c] * (nle[c] - ple[c] - 1); tempRes > res {
				res = tempRes
			}
		}
	}
	return res
}
