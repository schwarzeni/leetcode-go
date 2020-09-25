package _0205

// 2020.09.25
// https://leetcode-cn.com/problems/isomorphic-strings/
func isIsomorphic(s string, t string) bool {
	mapper := make(map[uint8]uint8)
	used := make(map[uint8]struct{})
	for i := range s {
		v, ok := mapper[s[i]]
		if !ok {
			if _, ok := used[t[i]]; ok {
				return false
			}
			mapper[s[i]] = t[i]
			used[t[i]] = struct{}{}
			continue
		}
		if v != t[i] {
			return false
		}
	}
	return true
}
