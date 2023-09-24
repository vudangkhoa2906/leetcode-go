package maximumlengthofrepeatedsubarray

import "testing"

func TestMaximumLengthOfRepeatedSubarray(t *testing.T) {
	nums1 := []int{1, 2, 3, 2, 1}
	nums2 := []int{3, 2, 1, 4, 7}
	if res := findLength(nums1, nums2); res != 3 {
		t.Errorf("Output: %d: Expected: %d\n", res, 3)
	}
}
