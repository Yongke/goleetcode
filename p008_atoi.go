package goleetcode

import (
	"strings"
	"math"
)

const (
	s_minus = 45
	s_add   = 43
	zero    = 48
	nine    = 57
)

func myAtoi(str string) int {
	str = strings.TrimLeft(str, " ")
	l := len(str)
	var sign int64
	var t,x int64
	t = 0
	sign = 1

	for i := 0; i < l; i++ {
		b := str[i]
		if i == 0 && is_sign(b) {
			if b == s_add {
				continue
			}
			sign = -1
			continue
		}
		if !is_num(b) {
			return int(x)
		}
		t = t*10 + int64(b) - zero
		x = sign * t
		if x < math.MinInt32 {
			return math.MinInt32
		}
		if x > math.MaxInt32 {
			return math.MaxInt32
		}
	}

	return int(x)
}

func is_sign(b byte) bool {
	return b == s_add || b == s_minus
}

func is_num(b byte) bool {
	return b >= zero && b <= nine
}
