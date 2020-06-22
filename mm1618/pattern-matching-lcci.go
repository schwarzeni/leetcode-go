package mm1618

// 2020.06.22
// https://leetcode-cn.com/problems/pattern-matching-lcci/
func patternMatching(pattern string, value string) bool {
	pMap := make(map[uint8]string)     // 记录 a,b 对应的字符串
	pvMap := make(map[string]struct{}) // 记录 a,b 对应字符串是否出现过

	var match func(pIdx, vIdx int) bool
	match = func(pIdx, vIdx int) bool {
		if len(pattern) == pIdx && len(value) == vIdx {
			return true
		}
		if len(pattern) == pIdx && len(value) > vIdx {
			return false
		}
		if v, ok := pMap[pattern[pIdx]]; ok {
			if vIdx+len(v) > len(value) || value[vIdx:vIdx+len(v)] != v {
				return false
			}
			return match(pIdx+1, vIdx+len(v))
		}
		for i := vIdx; i <= len(value); i++ {
			if _, ok := pvMap[value[vIdx:i]]; !ok {
				pvMap[value[vIdx:i]] = struct{}{}
				pMap[pattern[pIdx]] = value[vIdx:i]
				if match(pIdx+1, i) {
					return true
				}
				delete(pvMap, value[vIdx:i])
				delete(pMap, pattern[pIdx])
			}
		}
		return false
	}

	return match(0, 0)
}
