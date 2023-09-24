package majorityelement

import "testing"

func TestMajorityElement(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	if res := majorityElement(nums); res != 2 {
		t.Errorf("Output: %d: Expected: %v\n", res, 2)
	}
}
