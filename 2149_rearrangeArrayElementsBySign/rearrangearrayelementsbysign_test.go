package rearrangearrayelementsbysign

func rearrangeArray(nums []int) []int {
	res := make([]int, len(nums))
	idx1, idx2 := 0, 1
	for _, val := range nums {
		if val > 0 {
			res[idx1] = val
			idx1 += 2
		} else {
			res[idx2] = val
			idx2 += 2
		}
	}
	return res
}
