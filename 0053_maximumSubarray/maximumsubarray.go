package maximumsubarray

func maxSubArray(nums []int) int {
	res, maxInc := nums[0], nums[0]
	for idx := 1; idx < len(nums); idx++ {
		if maxInc < 0 {
			maxInc = nums[idx]
		} else {
			maxInc += nums[idx]
		}
		if res < maxInc {
			res = maxInc
		}
	}
	return res
}
