package _0003

// 2020.04.18
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
//func lengthOfLongestSubstring(s string) int {
//	marked := make(map[byte]int)
//	idx := 0
//	max := 0
//	begin := 0
//
//	for idx < len(s) {
//		b := s[idx]
//		i, ok := marked[b]
//		if !ok {
//			if idx == len(s)-1 {
//				if l := idx - begin + 1; max < l {
//					max = l
//				}
//			}
//			marked[b] = idx
//		} else {
//			l := idx - begin + 1
//			if i >= begin {
//				l--
//			}
//			if max < l {
//				max = l
//			}
//			if i >= begin {
//				begin = i + 1
//			}
//			marked[b] = idx
//		}
//		idx += 1
//	}
//
//	return max
//}

// better logic
func lengthOfLongestSubstring(s string) int {
	marked := make(map[byte]int)
	max := 0
	begin := 0

	for idx := 0; idx < len(s); idx++ {
		b := s[idx]
		if i, ok := marked[b]; ok && begin <= i { // 判断已出现字符的 idx 是否在 begin 之后
			begin = i + 1
		}
		if l := idx - begin + 1; l > max {
			max = l
		}
		marked[b] = idx
	}

	return max
}
