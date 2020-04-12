package _468

import (
	"strings"
)

// 2020.04.12
// https://leetcode-cn.com/problems/validate-ip-address/
const (
	T_IPv4    = "IPv4"
	T_IPv6    = "IPv6"
	T_INVALID = "Neither"

	IPv4_Spliter = "."
	IPv6_Spliter = ":"
)

func validIPAddress(IP string) string {
	isIPv4 := strings.Index(IP, IPv4_Spliter)
	isIPv6 := strings.Index(IP, IPv6_Spliter)

	if isIPv4 != -1 && isIPv6 != -1 {
		return T_INVALID
	}
	if isIPv4 != -1 && validIPv4(IP) {
		return T_IPv4
	}
	if isIPv6 != -1 && validIPv6(IP) {
		return T_IPv6
	}
	return T_INVALID
}

func validIPv4(IP string) bool {
	subIPs := strings.Split(IP, IPv4_Spliter)
	if len(subIPs) != 4 { // 验证长度是否合法
		return false
	}
	for _, s := range subIPs {
		if len(s) == 0 { // 验证每一个子串不为空
			return false
		}
		val := 0
		for idx, b := range s {
			// fix test_1 : idx != len(s)-1
			if b < '0' || b > '9' || (b == '0' && val == 0 && idx != len(s)-1) { // 0-9 且 不以 0 开头
				return false
			}
			val = int(b-48) + val*10
			if val > 255 { // 范围在 0 - 255 之间
				return false
			}
		}
	}
	return true
}

func validIPv6(IP string) bool {
	subIPs := strings.Split(IP, IPv6_Spliter)
	if len(subIPs) != 8 { // 验证长度是否合法
		return false
	}
	for _, s := range subIPs {
		if len(s) == 0 { // 验证每一个子串不为空
			return false
		}
		count := 0
		for _, b := range s {
			if (b >= 'a' && b <= 'f') || (b >= 'A' && b <= 'F') || (b >= '0' && b <= '9') {
				count += 1
				if count > 4 {
					return false
				}
				continue
			}
			return false
		}
	}
	return true
}

// A
// B
// C
// D
// E
// F
