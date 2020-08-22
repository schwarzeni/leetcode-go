package _1529

// 2020.08.22
// https://leetcode-cn.com/problems/bulb-switcher-iv/
func minFlips(target string) int {
	change, step := false, 0
	for i := range target {
		if (target[i] == '0' && change) || (target[i] == '1' && !change) {
			step++
			change = !change
		}
	}
	return step
}
