package _0028

// 2020.04.26
// https://leetcode-cn.com/problems/implement-strstr/
//func strStr(haystack string, needle string) int {
//	if len(needle) == 0 { // fix test_1
//		return 0
//	}
//	var cache = make([]int, len(needle))
//
//	count := 0
//	pre := ' '
//	for idx, v := range needle {
//		if idx != 0 && pre != v {
//			//pre = v // fix test_2
//			//count = 0
//			break
//		}
//		if idx == 0 {
//			pre = v
//		}
//		cache[idx] = count
//		count++
//	}
//
//	jumpIdx := 0
//LOOP:
//	for i := range haystack {
//		hi := i + jumpIdx
//		j := jumpIdx
//		if jumpIdx > 0 { // fix test_3
//			hi--
//			j--
//		}
//		for ni, nv := range needle[j:] {
//			if hi >= len(haystack) || haystack[hi] != uint8(nv) {
//				continue LOOP
//			}
//			jumpIdx = cache[ni+j]
//			hi++
//		}
//		return i
//	}
//
//	return -1
//}

func strStr(haystack string, needle string) int {
	ln := len(needle)
	if ln == 0 { // fix test_1
		return 0
	}

	for i := range haystack {
		if i+ln <= len(haystack) && haystack[i:i+ln] == needle {
			return i
		}
	}

	return -1
}
