package jumpgame

func canJump(nums []int) bool {
	len := len(nums)
	res := make([]bool, len)
	res[0] = true
	for i := 1; i < len; i++ {
		for j := 0; j < i; j++ {
			if i-j <= nums[j] && res[j] {
				res[i] = true
				break
			}
		}
	}
	return res[len-1]
}
