package maximumcandiesallocatedtokchildren

import "testing"

func TestMaximumCandiesAllocatedToKChildren(t *testing.T) {
	candies := []int{4, 7, 5}
	k := int64(4)
	if res := maximumCandies(candies, k); res != 3 {
		t.Errorf("Output: %d: Expected: %d\n", res, 3)
	}
}
