package courseschedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	for idx := range prerequisites {
		adj[prerequisites[idx][1]] = append(adj[prerequisites[idx][1]], prerequisites[idx][0])
	}

	marked, pathMarked := make([]bool, numCourses), make([]bool, numCourses)

	var depthFirstSearch func(idx int) bool
	depthFirstSearch = func(idx int) bool {
		if pathMarked[idx] {
			return true
		}
		marked[idx], pathMarked[idx] = true, true
		for _, otherIdx := range adj[idx] {
			if !marked[otherIdx] && depthFirstSearch(otherIdx) {
				return true
			} else if pathMarked[otherIdx] {
				return true
			}
		}
		pathMarked[idx] = false
		return false
	}

	for idx := 0; idx < numCourses; idx++ {
		if !marked[idx] && depthFirstSearch(idx) {
			return false
		}
	}
	return true
}
