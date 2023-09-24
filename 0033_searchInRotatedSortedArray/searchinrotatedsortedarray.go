package searchinrotatedsortedarray

func search(nums []int, target int) int {
	len := len(nums)
	lowIdx, highIdx := 0, len-1
	var midIdx int
	for lowIdx <= highIdx {
		midIdx = (lowIdx + highIdx) / 2
		if nums[midIdx] == target {
			return midIdx
		} else if nums[lowIdx] > nums[midIdx] { // second half is sorted
			if target >= nums[midIdx] && target <= nums[highIdx] { // if target is in second half
				lowIdx = midIdx + 1
			} else {
				highIdx = midIdx - 1
			}
		} else { // first half is sorted
			if target <= nums[midIdx] && target >= nums[lowIdx] { // if target is in first half
				highIdx = midIdx - 1
			} else {
				lowIdx = midIdx + 1
			}
		}
	}
	return -1
}
