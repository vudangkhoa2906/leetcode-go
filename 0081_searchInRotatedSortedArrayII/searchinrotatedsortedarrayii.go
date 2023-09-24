package searchinrotatedsortedarrayii

func search(nums []int, target int) bool {
	len := len(nums)
	lowIdx, highIdx := 0, len-1
	var midIdx int
	for lowIdx <= highIdx {
		midIdx = (lowIdx + highIdx) / 2
		if nums[midIdx] == target {
			return true
		} else if nums[lowIdx] == nums[midIdx] {
			lowIdx++
		} else if nums[highIdx] == nums[midIdx] {
			highIdx--
		} else if nums[lowIdx] > nums[midIdx] {
			if target >= nums[midIdx] && target <= nums[highIdx] {
				lowIdx = midIdx + 1
			} else {
				highIdx = midIdx - 1
			}
		} else {
			if target >= nums[lowIdx] && target <= nums[midIdx] {
				highIdx = midIdx - 1
			} else {
				lowIdx = midIdx + 1
			}
		}
	}
	return false
}
