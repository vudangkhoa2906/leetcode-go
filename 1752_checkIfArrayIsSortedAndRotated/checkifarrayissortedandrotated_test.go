package checkifarrayissortedandrotated

import "testing"

func TestCheckIfArrayIsSortedAndRotated(t *testing.T) {
	nums := []int{3, 4, 5, 1, 2}
	if res := check(nums); !res {
		t.Errorf("Output: %t: Expected: %t\n", res, true)
	}
}
