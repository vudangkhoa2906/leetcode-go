package missingnumber

import "testing"

func TestMissingNumber(t *testing.T) {
	nums := []int{3, 0, 1}
	if res := missingNumber(nums); res != 2 {
		t.Errorf("Output: %d: Expected: %d\n", res, 2)
	}
}
