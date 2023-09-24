package longestconsecutivesequence

func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})
	var count, res int
	var void struct{}
	for _, val := range nums {
		set[val] = void
	}
	for val := range set {
		if _, prs := set[val-1]; prs {
			continue
		} else {
			count = 1
			for i := 1; ; i++ {
				if _, prs = set[val+i]; prs {
					count++
				} else {
					if res < count {
						res = count
					}
					break
				}
			}
		}
	}
	return res
}
