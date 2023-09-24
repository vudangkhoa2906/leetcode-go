package sortcolors

func sortColors(nums []int) {
	len := len(nums)
	freq := make([]int, 2)
	for _, val := range nums {
		freq[val]++
	}
	for idx := 1; idx < 3; idx++ {
		freq[idx] += freq[idx-1]
	}
	tempRes := make([]int, len)
	for idx := len - 1; idx >= 0; idx-- {
		freq[nums[idx]]--
		tempRes[freq[nums[idx]]] = nums[idx]
	}
	copy(nums, tempRes)
}
