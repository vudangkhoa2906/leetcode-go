package common

func QuickSort(nums []int, lowIdx, highIdx int) {
	if lowIdx >= highIdx {
		return
	}
	pivotIdx := FindPivotIdx(nums, lowIdx, highIdx)
	QuickSort(nums, lowIdx, pivotIdx-1)
	QuickSort(nums, pivotIdx+1, highIdx)
}

func FindPivotIdx(nums []int, lowIdx, highIdx int) int {
	pivot := nums[lowIdx]
	leftIdx, rightIdx := lowIdx, highIdx
	for leftIdx < rightIdx {
		for nums[leftIdx] <= pivot && leftIdx <= highIdx {
			leftIdx++
		}
		for nums[rightIdx] > pivot && rightIdx >= lowIdx {
			rightIdx--
		}
		if leftIdx < rightIdx {
			nums[leftIdx], nums[rightIdx] = nums[rightIdx], nums[leftIdx]
		}
	}
	nums[lowIdx], nums[rightIdx] = nums[rightIdx], nums[lowIdx]
	return rightIdx
}
