package findfirstandlastpositionofelementinsortedarray

func searchRange(nums []int, target int) []int {
	len := len(nums)
	resLow, resHigh := -1, len
	lowIdx, highIdx := 0, len-1
	var midIdx int
	for lowIdx <= highIdx {
		midIdx = (lowIdx + highIdx) / 2
		if nums[midIdx] >= target {
			resLow, highIdx = midIdx, midIdx-1
		} else {
			lowIdx = midIdx + 1
		}
	}
	if resLow == -1 || nums[resLow] != target {
		return []int{-1, -1}
	}
	lowIdx, highIdx = 0, len-1
	for lowIdx <= highIdx {
		midIdx = (lowIdx + highIdx) / 2
		if nums[midIdx] > target {
			resHigh, highIdx = midIdx, midIdx-1
		} else {
			lowIdx = midIdx + 1
		}
	}
	return []int{resLow, resHigh - 1}
}
