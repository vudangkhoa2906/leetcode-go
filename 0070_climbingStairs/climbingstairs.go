package climbingstairs

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	resPrev, resPrevPrev := 1, 1
	var resCur int
	for i := 1; i < n; i++ {
		resCur = resPrev + resPrevPrev
		resPrev, resPrevPrev = resCur, resPrev
	}
	return resPrev
}
