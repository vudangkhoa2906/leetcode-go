package kokoeatingbananas

func minEatingSpeed(piles []int, h int) int {
	var highVal, lowVal int
	for _, val := range piles {
		if val > highVal {
			highVal = val
		}
	}
	var midVal, count int
	res := highVal
	for lowVal <= highVal {
		midVal = (lowVal + highVal) / 2
		if midVal == 0 {
			break
		}
		count = 0
		for _, val := range piles {
			if val%midVal != 0 {
				count += val/midVal + 1
			} else {
				count += val / midVal
			}
		}
		if count > h {
			lowVal = midVal + 1
		} else {
			if midVal < res {
				res = midVal
			}
			highVal = midVal - 1
		}
	}
	return res
}
