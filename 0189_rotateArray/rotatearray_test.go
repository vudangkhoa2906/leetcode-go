package rotatearray

import "testing"

func TestRotateArray(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	t.Logf("Output: %v\n", nums)
}
