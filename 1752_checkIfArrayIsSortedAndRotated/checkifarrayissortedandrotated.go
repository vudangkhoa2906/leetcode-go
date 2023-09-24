package checkifarrayissortedandrotated

func check(nums []int) bool {
	var rotated bool
	len := len(nums)
	for idx := 0; idx < len-1; idx++ {
		if !rotated {
			if nums[idx+1] < nums[idx] {
				rotated = true
			}
		} else {
			if nums[idx+1] < nums[idx] {
				return false
			}
		}
	}
	if rotated {
		return nums[0] >= nums[len-1]
	}
	return true
}
