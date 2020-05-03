package _0049

import "sort"

// 2020.05.03
// https://leetcode-cn.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]int)

	for idx, ss := range strs {
		s := []byte(ss)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		ss = string(s)
		a, ok := m[ss]
		if ok {
			m[ss] = append(a, idx)
		} else {
			m[ss] = []int{idx}
		}
	}

	var r [][]string
	for _, v := range m {
		rr := make([]string, len(v))
		for i, idx := range v {
			rr[i] = strs[idx]
		}
		r = append(r, rr)
	}
	return r
}
