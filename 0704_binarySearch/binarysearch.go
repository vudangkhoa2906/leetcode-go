package binarysearch

func search(nums []int, target int) int {
	lowIdx, highIdx := 0, len(nums)-1
	var midIdx int
	for highIdx >= lowIdx {
		midIdx = (lowIdx + highIdx) / 2
		if nums[midIdx] == target {
			return midIdx
		} else if nums[midIdx] > target {
			highIdx = midIdx - 1
		} else {
			lowIdx = midIdx + 1
		}
	}
	return -1
}
