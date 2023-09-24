package combinationsumii

import (
	"sort"
	"vudangkhoa2906-leetcode-go/common"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var bt []int
	var res [][]int
	nge := common.NextGreaterElement(candidates)

	var combinationSum2Rec func(idx, targetSum int)
	combinationSum2Rec = func(idx, targetSum int) {
		if idx == len(candidates) {
			if targetSum == 0 {
				tempRes := make([]int, len(bt))
				copy(tempRes, bt)
				res = append(res, tempRes)
			}
		} else {
			if targetSum >= candidates[idx] {
				bt = append(bt, candidates[idx])
				combinationSum2Rec(idx+1, targetSum-candidates[idx])
				bt = bt[:len(bt)-1]
			}
			upperBoundIdx := nge[idx]
			combinationSum2Rec(upperBoundIdx, targetSum)
		}
	}
	combinationSum2Rec(0, target)
	return res
}
