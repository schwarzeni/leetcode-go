package _0134

// 2020.09.17
// https://leetcode-cn.com/problems/gas-station/
func canCompleteCircuit(gas []int, cost []int) int {
	balance := 0
	for i := range gas {
		cost[i] = gas[i] - cost[i]
		balance += cost[i]
	}
	if balance < 0 {
		return -1
	}
	for i := 0; i < len(cost); {
		if cost[i] < 0 {
			i++
			continue
		}
		balance := 0
		j := i
		for {
			balance += cost[j]
			if balance < 0 {
				if j < i {
					return -1
				}
				i = j + 1
				break
			}
			j++
			if j >= len(cost) {
				j = 0
			}
			if j == i {
				return i
			}
		}
	}
	return -1
}

// [1,2,3,4,5]
// [3,4,5,1,2]
