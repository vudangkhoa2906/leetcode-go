package containerwithmostwater

import "vudangkhoa2906-leetcode-go/common"

func maxArea(height []int) int {
	var res, tempHeight int
	size := len(height)
	lowIdx, highIdx := 0, size-1
	for lowIdx < highIdx {
		if tempRes := (highIdx - lowIdx) * common.Min(height[lowIdx], height[highIdx]); tempRes > res {
			res = tempRes
		}
		if height[lowIdx] < height[highIdx] {
			tempHeight = height[lowIdx]
			for lowIdx < size && height[lowIdx] <= tempHeight {
				lowIdx++
			}
		} else {
			tempHeight = height[highIdx]
			for highIdx >= 0 && height[highIdx] <= tempHeight {
				highIdx--
			}
		}
	}
	return res
}
