package subsets

func subsets(nums []int) [][]int {
	var res [][]int
	var bt []int
	lenNums := len(nums)

	var subsetsRec func(idx int)
	subsetsRec = func(idx int) {
		if idx == lenNums {
			tempRes := make([]int, len(bt))
			copy(tempRes, bt)
			res = append(res, tempRes)
		} else {
			bt = append(bt, nums[idx])
			subsetsRec(idx + 1)
			bt = bt[:len(bt)-1]
			subsetsRec(idx + 1)
		}
	}

	subsetsRec(0)
	return res
}
