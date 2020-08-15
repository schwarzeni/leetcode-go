package _0331

import "strings"

// 2020.08.14
// https://leetcode-cn.com/problems/verify-preorder-serialization-of-a-binary-tree/
func isValidSerialization(preorder string) bool {
	slot := 1
	for _, s := range strings.Split(preorder, ",") {
		slot--
		if slot < 0 {
			return false
		}
		if s != "#" {
			slot += 2
		}
	}
	return slot == 0
}
