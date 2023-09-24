package nondecreasingsubsequences

func findSubsequences(nums []int) [][]int {
	lenNums := len(nums)
	var res [][]int
	var bt []int
	var tempRes []int

	var findSubsequencesRec func(idx int)
	findSubsequencesRec = func(idx int) {
		if idx == lenNums {
			if lenBt := len(bt); lenBt > 1 {
				tempRes = make([]int, lenBt)
				copy(tempRes, bt)
				res = append(res, tempRes)
			}
			return
		}
		if lenBt := len(bt); lenBt > 0 {
			if bt[lenBt-1] <= nums[idx] {
				bt = append(bt, nums[idx])
				findSubsequencesRec(idx + 1)
				bt = bt[:lenBt]
			}
		} else {
			bt = append(bt, nums[idx])
			findSubsequencesRec(idx + 1)
			bt = bt[:lenBt]
		}
		if lenBt := len(bt); lenBt > 0 {
			if bt[lenBt-1] != nums[idx] {
				findSubsequencesRec(idx + 1)
			}
		} else {
			findSubsequencesRec(idx + 1)
		}
	}

	findSubsequencesRec(0)
	return res
}
