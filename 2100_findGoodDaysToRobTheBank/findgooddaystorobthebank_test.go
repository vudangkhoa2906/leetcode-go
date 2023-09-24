package findgooddaystorobthebank

import "testing"

func TestFindGoodDaysToRobTheBank(t *testing.T) {
	security := []int{5, 3, 3, 3, 5, 6, 2}
	time := 2
	res := goodDaysToRobBank(security, time)
	t.Logf("Output: %v\n", res)
}
