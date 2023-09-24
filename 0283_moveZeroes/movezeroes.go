package movezeroes

func moveZeroes(nums []int) {
	var count int
	for idx := range nums {
		if nums[idx] != 0 {
			nums[idx], nums[count] = nums[count], nums[idx]
			count++
		}
	}
}
