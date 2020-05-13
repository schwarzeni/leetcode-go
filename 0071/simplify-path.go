package _0071

import "strings"

// 2020.05.13
func simplifyPath(path string) string {
	var p []string
	for _, s := range strings.Split(path, "/") {
		if s == "" || s == "." {
			continue
		} else if s == ".." {
			if len(p) > 0 {
				p = p[:len(p)-1]
			}
		} else {
			p = append(p, s)
		}
	}
	if len(p) == 0 {
		return "/"
	} else {
		return "/" + strings.Join(p, "/")
	}
}
