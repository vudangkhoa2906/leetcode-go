package jumpgameii

func jump(nums []int) int {
	len := len(nums)
	res := make([]int, len)
	res[0] = 0
	for i := 1; i < len; i++ {
		res[i] = -1
		for j := 0; j < i; j++ {
			if res[j] >= 0 && i-j <= nums[j] {
				res[i] = res[j] + 1
				break
			}
		}
	}
	return res[len-1]
}
