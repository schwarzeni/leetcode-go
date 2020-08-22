package _0165

import (
	"strconv"
	"strings"
)

// 2020.08.22
//  https://leetcode-cn.com/problems/compare-version-numbers/
func compareVersion(version1 string, version2 string) int {
	v1 := trim(toInt(version1))
	v2 := trim(toInt(version2))

	for len(v1) > 0 && len(v2) > 0 {
		if r := compare(v1[0], v2[0]); r != 0 {
			return r
		}
		v1 = v1[1:]
		v2 = v2[1:]
	}

	if len(v1) > 0 {
		return 1
	}
	if len(v2) > 0 {
		return -1
	}
	return 0
}

func toInt(vStr string) []int {
	vList := strings.Split(vStr, ".")
	res := make([]int, len(vList))
	for i, s := range vList {
		v, _ := strconv.Atoi(s)
		res[i] = v
	}
	return res
}

func trim(vList []int) []int {
	for i := len(vList) - 1; i >= 0; i-- {
		if vList[i] != 0 {
			break
		}
		vList = vList[:len(vList)-1]
	}
	if len(vList) == 0 {
		vList = []int{0}
	}
	return vList
}

func compare(v1, v2 int) int {
	switch {
	case v1 < v2:
		return -1
	case v1 > v2:
		return 1
	default:
		return 0
	}
}
