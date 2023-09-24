package subsetsii

import (
	"sort"
	"vudangkhoa2906-leetcode-go/common"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var bt []int
	nge := common.NextGreaterElement(nums)

	var subsetsWithDupRec func(idx int)
	subsetsWithDupRec = func(idx int) {
		if idx == len(nums) {
			tempRes := make([]int, len(bt))
			copy(tempRes, bt)
			res = append(res, tempRes)
		} else {
			bt = append(bt, nums[idx])
			subsetsWithDupRec(idx + 1)
			bt = bt[:len(bt)-1]
			subsetsWithDupRec(nge[idx])
		}
	}

	subsetsWithDupRec(0)
	return res
}
