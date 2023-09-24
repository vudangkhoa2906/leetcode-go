package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	res, curVal, curIdx := 1, nums[0], 0
	for idx := 1; idx < len(nums); idx++ {
		if nums[idx] > curVal {
			res++
			curIdx++
			nums[curIdx], curVal = nums[idx], nums[idx]
		}
	}
	return res
}
