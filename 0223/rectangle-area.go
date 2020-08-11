package _0223

// 2020.08.11
// https://leetcode-cn.com/problems/rectangle-area/
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	s1 := S(A, B, C, D)
	s2 := S(E, F, G, H)
	w, h := 0, 0
	if C-A > G-E {
		w = L(E, G, A, C)
	} else {
		w = L(A, C, E, G)
	}
	if D-B > H-F {
		h = L(F, H, B, D)
	} else {
		h = L(B, D, F, H)
	}
	return s1 + s2 - w*h
}

// x1r - x1l <= x2r - x2l
func L(x1l, x1r, x2l, x2r int) int {
	if x1r <= x2l || x1l >= x2r {
		return 0
	}
	if x1l >= x2l && x1r <= x2r {
		return x1r - x1l
	}
	if x1l <= x2l {
		return x1r - x2l
	} else {
		return x2r - x1l
	}
}

func S(x1, y1, x2, y2 int) int {
	return abs(x2-x1) * (y2 - y1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 6*4 + 9*3 - 3*2
// 24 + 27 - 6
// 24 + 21 = 45
