package coursescheduleii

func findOrder(numCourses int, prerequisites [][]int) []int {
	prerequisitesTemp := make([][]int, numCourses)
	for idx := range prerequisites {
		prerequisitesTemp[prerequisites[idx][0]] = append(prerequisitesTemp[prerequisites[idx][0]], prerequisites[idx][1])
	}
	canBeFinished := make([]int, numCourses)
	res := make([]int, 0, numCourses)

	var canFinishRec func(idx int) int
	canFinishRec = func(idx int) int {
		if canBeFinished[idx] != 0 {
			return canBeFinished[idx]
		}
		if len(prerequisitesTemp[idx]) == 0 {
			canBeFinished[idx] = 1
			res = append(res, idx)
			return canBeFinished[idx]
		}
		canBeFinished[idx] = -1
		for _, otherIdx := range prerequisitesTemp[idx] {
			if canFinishRec(otherIdx) == -1 {
				return canBeFinished[idx]
			}
		}
		canBeFinished[idx] = 1
		res = append(res, idx)
		return canBeFinished[idx]
	}

	for idx := 0; idx < numCourses; idx++ {
		if canFinishRec(idx) == -1 {
			return []int{}
		}
	}
	return res
}
