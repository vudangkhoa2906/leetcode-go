package frequencyofthemostfrequentelement

import (
	"math"
	"sort"
)

func maxFrequency(nums []int, k int) int {
	res, len := math.MinInt, len(nums)
	var target, tempRes, idx int
	sort.Ints(nums)
	for highIdx := len - 1; highIdx >= 0; {
		tempRes, target = 0, k
		for lowIdx := highIdx; lowIdx >= 0; lowIdx-- {
			target -= nums[highIdx] - nums[lowIdx]
			if target < 0 {
				if tempRes > res {
					res = tempRes
				}
				break
			}
			tempRes++
			if lowIdx == 0 {
				if tempRes > res {
					res = tempRes
				}
				break
			}
		}
		for idx = highIdx - 1; idx >= 0; idx-- {
			if nums[idx] < nums[highIdx] {
				break
			}
		}
		highIdx = idx
	}
	return res
}
