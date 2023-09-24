package singlenumber

import "testing"

func TestSingleNumber(t *testing.T) {
	nums := []int{4, 1, 2, 1, 2}
	if res := singleNumber(nums); res != 4 {
		t.Errorf("Output: %d: Expected: %d\n", res, 4)
	}
}
