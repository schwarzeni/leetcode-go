package _0014

// 2020.04.20
// https://leetcode-cn.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	var comm []byte

	if len(strs) == 0 {
		return ""
	}

LOOP:
	for idx, m := range strs[0] {
		mm := uint8(m)
		for _, s := range strs[1:] {
			if len(s) <= idx || s[idx] != mm {
				break LOOP
			}
		}
		comm = append(comm, mm)
	}

	return string(comm)
}
