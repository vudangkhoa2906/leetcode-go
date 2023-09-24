package common

func MergeSort(nums []int, lowIdx, highIdx int) {
	if lowIdx >= highIdx {
		return
	}
	midIdx := (lowIdx + highIdx) / 2
	MergeSort(nums, lowIdx, midIdx)
	MergeSort(nums, midIdx+1, highIdx)
	Merge(nums, lowIdx, midIdx, highIdx)
}

func Merge(nums []int, lowIdx, midIdx, highIdx int) {
	var tempRes []int
	leftIdx, rightIdx := lowIdx, midIdx+1
	for leftIdx <= midIdx && rightIdx <= highIdx {
		if nums[leftIdx] <= nums[rightIdx] {
			tempRes = append(tempRes, nums[leftIdx])
			leftIdx++
		} else {
			tempRes = append(tempRes, nums[rightIdx])
			rightIdx++
		}
	}
	for ; leftIdx <= midIdx; leftIdx++ {
		tempRes = append(tempRes, nums[leftIdx])
	}
	for ; rightIdx <= highIdx; rightIdx++ {
		tempRes = append(tempRes, nums[rightIdx])
	}
	copy(nums[lowIdx:highIdx+1], tempRes)
}
