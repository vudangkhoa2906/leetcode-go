package permutations

func permute(nums []int) [][]int {
	lenNums := len(nums)
	marked := make([]bool, lenNums)
	var bt []int
	var res [][]int

	var permuteRec func()
	permuteRec = func() {
		if lenBt := len(bt); lenBt == lenNums {
			tempRes := make([]int, lenBt)
			copy(tempRes, bt)
			res = append(res, tempRes)
		} else {
			for idx, val := range nums {
				if !marked[idx] {
					marked[idx] = true
					bt = append(bt, val)
					permuteRec()
					bt = bt[:len(bt)-1]
					marked[idx] = false
				}
			}
		}
	}

	permuteRec()
	return res
}
