package trappingrainwater

import "vudangkhoa2906-leetcode-go/common"

func trap(height []int) int {
	var res, lowIdx, highIdx int
	lenHeight := len(height)
	nge, pge := common.NextGreaterElement(height), common.PreviousGreaterElement(height)
	for idx := 1; idx < lenHeight-1; idx++ {
		if nge[idx] != lenHeight && pge[idx] != -1 {
			lowIdx, highIdx = pge[idx], nge[idx]
			for pge[lowIdx] != -1 {
				lowIdx = pge[lowIdx]
			}
			for nge[highIdx] != lenHeight {
				highIdx = nge[highIdx]
			}
			res += common.Min(height[lowIdx], height[highIdx]) - height[idx]
		}
	}
	return res
}
