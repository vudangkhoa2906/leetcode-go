package maxconsecutiveones

func findMaxConsecutiveOnes(nums []int) int {
	var res, count int
	for _, val := range nums {
		if val == 1 {
			count++
		} else {
			if count > res {
				res = count
			}
			count = 0
		}
	}
	if count > res {
		res = count
	}
	return res
}
