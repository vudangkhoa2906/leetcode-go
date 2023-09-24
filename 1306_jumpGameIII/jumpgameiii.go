package jumpgameiii

import "fmt"

func canReach(arr []int, start int) bool {
	lenArr := len(arr)
	adj := make([][]int, lenArr)
	for idx := 0; idx < lenArr; idx++ {
		if lowIdx := idx - arr[idx]; lowIdx >= 0 {
			adj[idx] = append(adj[idx], lowIdx)
		}
		if highIdx := idx + arr[idx]; highIdx < lenArr {
			adj[idx] = append(adj[idx], highIdx)
		}
	}

	fmt.Println(adj)

	marked := make([]bool, lenArr)

	var depthFirstSearch func(idx int)
	depthFirstSearch = func(idx int) {
		if !marked[idx] {
			marked[idx] = true
			for _, otherIdx := range adj[idx] {
				depthFirstSearch(otherIdx)
			}
		}
	}

	depthFirstSearch(start)
	for idx := 0; idx < lenArr; idx++ {
		if arr[idx] == 0 && !marked[idx] {
			return false
		}
	}
	return true
}
