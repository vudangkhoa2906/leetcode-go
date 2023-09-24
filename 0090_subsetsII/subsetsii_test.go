package subsetsii

import "testing"

func TestSubsetsII(t *testing.T) {
	nums := []int{1, 2, 2}
	res := subsetsWithDup(nums)
	t.Logf("Output: %v\n", res)
}
