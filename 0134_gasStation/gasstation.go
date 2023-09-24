package gasstation

func canCompleteCircuit(gas []int, cost []int) int {
	var remain, idxOther, delta int
	lenGas := len(gas)
	for idx := 0; idx < lenGas; idx = idxOther + 1 {
		remain = 0
		for delta = 0; delta < lenGas; delta++ {
			if idxOther = idx + delta; idxOther >= lenGas {
				idxOther -= lenGas
			}
			remain += gas[idxOther] - cost[idxOther]
			if remain < 0 {
				break
			}
		}
		if remain >= 0 {
			return idx
		}
		if idx+delta >= lenGas {
			break
		}
	}
	return -1
}
