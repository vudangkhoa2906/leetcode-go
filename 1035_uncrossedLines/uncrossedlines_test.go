package uncrossedlines

import "testing"

func TestUncrossedLines(t *testing.T) {
	nums1, nums2 := []int{1, 4, 2}, []int{1, 2, 4}
	if res := maxUncrossedLines(nums1, nums2); res != 2 {
		t.Errorf("Output: %d: Expected: %d\n", res, 2)
	}
}
