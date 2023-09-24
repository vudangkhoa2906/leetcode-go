package missingnumber

func missingNumber(nums []int) int {
	len := len(nums)
	var sum int
	for _, val := range nums {
		sum += val
	}
	return len*(len+1)/2 - sum
}
