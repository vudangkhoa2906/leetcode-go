package largestrectangleinhistogram

import (
	"vudangkhoa2906-leetcode-go/common"
)

func largestRectangleArea(heights []int) int {
	size := len(heights)
	var stack []common.SliceElem
	nle, ple := make([]int, size), make([]int, size)

	/* next lesser element */
	for idx := size - 1; idx >= 0; idx-- {
		for count := len(stack); count > 0 && stack[count-1].Val >= heights[idx]; count = len(stack) {
			stack = stack[:count-1]
		}
		if count := len(stack); count > 0 {
			nle[idx] = stack[count-1].Idx
		} else {
			nle[idx] = size
		}
		stack = append(stack, common.SliceElem{Idx: idx, Val: heights[idx]})
	}
	stack = stack[:0]

	/* previous lesser element */
	for idx := 0; idx < size; idx++ {
		for count := len(stack); count > 0 && stack[count-1].Val >= heights[idx]; count = len(stack) {
			stack = stack[:count-1]
		}
		if count := len(stack); count > 0 {
			ple[idx] = stack[count-1].Idx
		} else {
			ple[idx] = -1
		}
		stack = append(stack, common.SliceElem{Idx: idx, Val: heights[idx]})
	}

	var tempRes, res int
	for idx := 0; idx < size; idx++ {
		if tempRes = heights[idx] * (nle[idx] - ple[idx] - 1); tempRes > res {
			res = tempRes
		}
	}

	return res
}
