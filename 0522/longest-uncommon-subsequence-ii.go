package _0522

// 2020.05.22
// https://leetcode-cn.com/problems/longest-uncommon-subsequence-ii/
func findLUSlength(strs []string) int {
	maxL := -1
	cache := make(map[string]struct{})
LOOP:
	for idx, s := range strs {
		if _, ok := cache[s]; ok {
			continue
		}
		cache[s] = struct{}{}
		for i, ss := range strs {
			if i != idx && len(ss) >= len(s) {
				if contains(ss, s) {
					continue LOOP
				}
			}
		}
		if maxL < len(s) {
			maxL = len(s)
		}
	}
	return maxL
}

func contains(ss, s string) bool {
	i, j := 0, 0
	if len(ss) < len(s) {
		return false
	}
	if len(ss) == len(s) {
		return ss == s
	}
	for i < len(ss) && j < len(s) {
		if ss[i] == s[j] {
			i++
			j++
		} else {
			i++
		}
	}
	return j >= len(s)
}
