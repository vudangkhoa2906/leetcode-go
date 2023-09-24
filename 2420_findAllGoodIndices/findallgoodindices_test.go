package findallgoodindices

import "testing"

func TestFindAllGoodIndices(t *testing.T) {
	nums := []int{878724, 201541, 179099, 98437, 35765, 327555, 475851, 598885, 849470, 943442}
	k := 4
	res := goodIndices(nums, k)
	t.Logf("Output: %v\n", res)
}
