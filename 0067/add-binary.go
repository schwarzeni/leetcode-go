package _0067

// 2020.05.11
// https://leetcode-cn.com/problems/add-binary/
func addBinary(a string, b string) string {
	idxa, idxb, idxr, result := len(a)-1, len(b)-1, 0, []uint8{}
	if len(a) > len(b) {
		result, idxr = make([]uint8, len(a)+1), len(a)
	} else {
		result, idxr = make([]uint8, len(b)+1), len(b)
	}
	result[0] = '0'

	var pre uint8 = '0'
	for idxa >= 0 && idxb >= 0 {
		result[idxr], pre = cal(a[idxa], b[idxb], pre)
		idxr--
		idxa--
		idxb--
	}
	for idxa >= 0 {
		result[idxr], pre = cal(a[idxa], pre)
		idxr--
		idxa--
	}
	for idxb >= 0 {
		result[idxr], pre = cal(b[idxb], pre)
		idxr--
		idxb--
	}
	if pre == '1' {
		result[0] = '1'
		return string(result)
	}
	return string(result[1:])

}

func cal(d ...uint8) (curr, pre uint8) {
	count1 := 0
	for _, v := range d {
		if v == '1' {
			count1++
		}
	}
	if count1 == 3 {
		return '1', '1'
	} else if count1 == 2 {
		return '0', '1'
	} else if count1 == 1 {
		return '1', '0'
	} else {
		return '0', '0'
	}
}
