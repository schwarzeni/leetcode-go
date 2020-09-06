package _0179

import (
	"sort"
	"strconv"
	"strings"
)

// 2020.09.06
// https://leetcode-cn.com/problems/largest-number/
func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	strs := make([]string, len(nums))
	for i := range nums {
		strs[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(strs, func(i, j int) bool { return strs[i]+strs[j] > strs[j]+strs[i] })
	if res := strings.TrimLeft(strings.Join(strs, ""), "0"); len(res) != 0 {
		return res
	}
	return "0"
}

// sort.Slice(strs, func(i, j int) bool {
//     str1 := strs[i] + strs[j]
//     str2 := strs[j] + strs[i]
//     i, j = 0, 0
//     for i < len(str1) {
//         if str1[i] == str2[j] {
//             i ++
//             j ++
//             continue
//         }
//         return str1[i] > str2[j]
//     }
//     return false
// })
