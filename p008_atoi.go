package goleetcode

import (
	"math"
	"strings"
)

const (
	s_minus = 45
	s_add   = 43
	zero    = 48
	nine    = 57
)

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	l := len(str)
	is_neg := false
	var t uint64
	t = 0

	for i := 0; i < l; i++ {
		b := str[i]
		if i == 0 && is_sign(b) {
			if b == s_add {
				continue
			}
			is_neg = true
			continue
		}
		if !is_num(b) {
			return simple_ret(is_neg, t)
		}
		t = t*10 + uint64(b) - zero
		if is_neg && t > math.MaxInt32 + 1 {
			return math.MinInt32
		}
		if !is_neg &&  t > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return simple_ret(is_neg, t)
}

func is_sign(b byte) bool {
	return b == s_add || b == s_minus
}

func is_num(b byte) bool {
	return b >= zero && b <= nine
}

func simple_ret(is_neg bool, t uint64) int {
	if is_neg {
		return -int(t)
	}

	return int(t)
}
