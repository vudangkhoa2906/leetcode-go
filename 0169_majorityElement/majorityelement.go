package majorityelement

func majorityElement(nums []int) int {
	res, count := nums[0], 1
	for idx := 1; idx < len(nums); idx++ {
		if count == 0 {
			res, count = nums[idx], 1
			continue
		}
		if nums[idx] != res {
			count--
		} else {
			count++
		}
	}
	return res
}
