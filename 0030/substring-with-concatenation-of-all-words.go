package _0030

// 2020.04.27
// https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/

// this version 700ms
//func findSubstring(s string, words []string) []int {
//	if len(words) == 0 {
//		return nil
//	}
//	var (
//		size     = len(words[0])
//		result   []int
//		idx      = 0
//		wordsMap map[string]int
//		resetMap = func() {
//			wordsMap = make(map[string]int)
//			for _, w := range words {
//				wordsMap[w] += 1
//			}
//		}
//	)
//
//	resetMap()
//
//	for idx <= len(s)-size {
//		if _, ok := wordsMap[s[idx:idx+size]]; ok {
//			i := idx
//			count := len(words)
//			for i <= len(s)-size {
//				if v, ok := wordsMap[s[i:i+size]]; !ok || v == 0 {
//					break
//				}
//				wordsMap[s[i:i+size]] -= 1
//				count--
//				if count == 0 {
//					result = append(result, idx)
//					break
//				}
//				i += size
//			}
//			resetMap()
//		}
//		idx++
//	}
//
//	return result
//}

// this version 600ms
//func findSubstring(s string, words []string) []int {
//	if len(words) == 0 {
//		return nil
//	}
//	var (
//		size     = len(words[0])
//		result   []int
//		idx      = 0
//		wordsMap map[string]int
//		resetMap = func() {
//			wordsMap = make(map[string]int)
//			for _, w := range words {
//				wordsMap[w] += 1
//			}
//		}
//	)
//
//	if size == 0 {
//		result = make([]int, len(s)+1)
//		for idx := range result {
//			result[idx] = idx
//		}
//		return result
//	}
//
//	resetMap()
//
//	for idx < size {
//		l := idx
//		r := idx
//		count := len(words)
//
//		for r <= len(s)-size {
//			v, ok := wordsMap[s[r:r+size]]
//			if !ok {
//				r += size
//				l = r
//				count = len(words)
//				resetMap()
//			} else if v == 0 {
//				for l <= r {
//					wordsMap[s[l:l+size]] += 1
//					count++
//					if s[l:l+size] == s[r:r+size] {
//						break
//					}
//					l += size
//				}
//				l += size
//			} else {
//				wordsMap[s[r:r+size]] -= 1
//				count--
//				if count == 0 {
//					result = append(result, l)
//					count++
//					wordsMap[s[l:l+size]] += 1
//					l += size
//				}
//				r += size
//			}
//		}
//
//		resetMap()
//		idx++
//	}
//
//	return result
//}

// this is the best version
// use two map
// 8ms
func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return nil
	}
	var (
		size     = len(words[0])
		result   []int
		idx      = 0
		wordsMap map[string]int
	)

	if size == 0 {
		result = make([]int, len(s)+1)
		for idx := range result {
			result[idx] = idx
		}
		return result
	}

	wordsMap = make(map[string]int)
	for _, w := range words {
		wordsMap[w] += 1
	}

	for idx < size {
		l := idx
		r := idx
		count := len(words)
		slideMap := make(map[string]int) // 用于计数，省去直接操作 wordsMap 后重置所花费的时间

		for r <= len(s)-size {
			sCount, ok := wordsMap[s[r:r+size]]
			v, _ := slideMap[s[r:r+size]]
			if !ok { // 不存在
				r += size
				l = r
				count = len(words)
				slideMap = make(map[string]int)
			} else if v >= sCount { // 出现次数已达上届
				for l <= r {
					slideMap[s[l:l+size]] -= 1
					count++
					if s[l:l+size] == s[r:r+size] {
						break
					}
					l += size
				}
				l += size
			} else { // 其他情况
				slideMap[s[r:r+size]] += 1
				count--
				if count == 0 { // 统计到一个结果
					result = append(result, l)
					count++
					slideMap[s[l:l+size]] -= 1
					l += size
				}
				r += size
			}
		}
		idx++
	}

	return result
}
