package _1016

import (
	"strconv"
	"strings"
)

// 2020.05.25
// https://leetcode-cn.com/problems/binary-string-with-substrings-representing-1-to-n/
func queryString(S string, N int) bool {
	for i := N; i > 0; i-- {
		if !strings.Contains(S, strconv.FormatInt(int64(i), 2)) {
			return false
		}
	}
	return true
}

func toBinaryString(n int) string {
	r := []byte{}
	for n > 0 {
		if n%2 == 0 {
			r = append(r, '0')
		} else {
			r = append(r, '1')
		}
		n = n / 2
	}
	i, j := 0, len(r)-1
	for i < j {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
	return string(r)
}
