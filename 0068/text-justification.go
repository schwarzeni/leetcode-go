package _0068

import "strings"

// 2020.05.11
// https://leetcode-cn.com/problems/text-justification/
func fullJustify(words []string, maxWidth int) []string {
	widx, count, sidx, ans := 0, 0, 0, []string{}
	for {
		if widx >= len(words) || count == maxWidth+1 || count+len(words[widx]) > maxWidth {
			mode := 0
			if widx >= len(words) || sidx+1 == widx {
				mode = 1
			}
			s := lineJustify(words[sidx:widx], maxWidth, mode)
			ans = append(ans, s)
			if widx >= len(words) {
				return ans
			}
			count = 0
			sidx = widx
			continue
		}
		count += 1 + len(words[widx])
		widx++
	}
}

// mode: default: 0, left: 1
func lineJustify(words []string, maxWidth int, mode int) string {
	rs := strings.Builder{}
	wordsLen := 0
	for _, w := range words {
		wordsLen += len(w)
	}
	switch mode {
	case 1:
		for _, w := range words[:len(words)-1] {
			rs.WriteString(w)
			rs.WriteByte(' ')
		}
		rs.WriteString(words[len(words)-1])
		rs.WriteString(strings.Repeat(" ", maxWidth-wordsLen-(len(words)-1)))
	default:
		// 插入空格算法：
		// 空格总长度为8，插入3个空格
		// 8 3 ==> [ 0 0 2 ]
		// 6 2 ==> [ 0 3 2 ]
		// 3 1 ==> [ 3 3 2 ]
		spaceLs := make([]int, len(words)-1) // 记录插入空格的长度
		spaceWidth, idx := maxWidth-wordsLen, len(words)-1
		for idx > 0 {
			l := spaceWidth / idx
			spaceLs[idx-1] = l
			spaceWidth = spaceWidth - l
			idx--
		}
		for idx, w := range words[:len(words)-1] {
			rs.WriteString(w)
			rs.WriteString(strings.Repeat(" ", spaceLs[idx]))
		}
		rs.WriteString(words[len(words)-1])
	}
	return rs.String()
}
