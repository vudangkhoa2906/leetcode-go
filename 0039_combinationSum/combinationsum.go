package combinationsum

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var bt []int

	var combinationSumRec func(idx, targetSum int)
	combinationSumRec = func(idx, targetSum int) {
		if idx == len(candidates) {
			if targetSum == 0 {
				tempRes := make([]int, len(bt))
				copy(tempRes, bt)
				res = append(res, tempRes)
			}
		} else {
			if targetSum >= candidates[idx] {
				bt = append(bt, candidates[idx])
				combinationSumRec(idx, targetSum-candidates[idx])
				bt = bt[:len(bt)-1]
			}
			combinationSumRec(idx+1, targetSum)
		}
	}
	combinationSumRec(0, target)
	return res
}
