package rotatearray

func rotate(nums []int, k int) {
	len := len(nums)
	if len == 0 {
		return
	}
	k = k % len
	copy(nums, append(nums[len-k:], nums[:len-k]...))
}
