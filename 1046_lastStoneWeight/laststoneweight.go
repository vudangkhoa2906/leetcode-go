package laststoneweight

import "vudangkhoa2906-leetcode-go/common"

func lastStoneWeight(stones []int) int {
	common.Heapify(stones)
	var stone1, stone2 int
	for count := len(stones); count > 1; count = len(stones) {
		stone1, stones = common.DelMax(stones)
		stone2, stones = common.DelMax(stones)
		if stone := stone1 - stone2; stone > 0 {
			stones = append(stones, stone)
			common.Heapify(stones)
		}
	}
	if len(stones) == 0 {
		return 0
	} else {
		return stones[0]
	}
}
