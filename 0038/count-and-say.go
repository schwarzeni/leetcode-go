package _0038

import "strings"

// 2020.04.29
// https://leetcode-cn.com/problems/count-and-say/
func countAndSay(n int) string {
	pre := "1"
	for i := 2; i <= n; i++ {
		cc := int32(pre[0])
		count := 0
		sb := strings.Builder{}
		for _, c := range pre {
			if c == cc {
				count++
			} else {
				sb.WriteByte(byte(count + 48))
				sb.WriteByte(byte(cc))
				cc = c
				count = 1
			}
		}
		sb.WriteByte(byte(count + 48))
		sb.WriteByte(byte(cc))
		pre = sb.String()
	}

	return pre
}
