package findpeakelement

func findPeakElement(nums []int) int {
	len := len(nums)
	if len == 1 || nums[0] > nums[1] {
		return 0
	} else if nums[len-1] > nums[len-2] {
		return len - 1
	}
	lowIdx, highIdx := 1, len-2
	var midIdx int
	for lowIdx <= highIdx {
		midIdx = (lowIdx + highIdx) / 2
		if nums[midIdx] > nums[midIdx-1] && nums[midIdx] > nums[midIdx+1] {
			return midIdx
		} else if nums[midIdx] > nums[midIdx-1] {
			lowIdx = midIdx + 1
		} else {
			highIdx = midIdx - 1
		}
	}
	return midIdx
}
