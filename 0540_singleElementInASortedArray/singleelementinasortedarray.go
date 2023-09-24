package singleelementinasortedarray

func singleNonDuplicate(nums []int) int {
	lowIdx, highIdx := 0, len(nums)-1
	var midIdx int
	for lowIdx <= highIdx {
		midIdx = (lowIdx + highIdx) / 2
		if lowIdx == highIdx || nums[midIdx] != nums[midIdx-1] && nums[midIdx] != nums[midIdx+1] {
			return nums[midIdx]
		} else if (highIdx-lowIdx)%4 == 0 {
			if nums[midIdx] == nums[midIdx-1] {
				highIdx = midIdx - 2
			} else {
				lowIdx = midIdx + 2
			}
		} else {
			if nums[midIdx] == nums[midIdx-1] {
				lowIdx = midIdx + 1
			} else {
				highIdx = midIdx - 1
			}
		}
	}
	return -1
}
