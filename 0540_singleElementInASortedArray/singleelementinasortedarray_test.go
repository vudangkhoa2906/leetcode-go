package singleelementinasortedarray

import "testing"

func TestSingleElementInASortedArray(t *testing.T) {
	nums := []int{3, 3, 7, 7, 10, 11, 11}
	if res := singleNonDuplicate(nums); res != 10 {
		t.Errorf("Output: %d: Expected: %d\n", res, 2)
	}
}
