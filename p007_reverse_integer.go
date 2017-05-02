package goleetcode

import "math"

func reverse(x int) int {
	if x == 0 {
		return 0
	}

	var u uint64
	sign := false
	if x < 0 {
		sign = true
		u = uint64(-x)
	} else {
		u = uint64(x)
	}

	t := reverse_aux(u)
	if t > math.MaxInt32 {
		return 0
	}

	if sign {
		return -int(t)
	}
	return int(t)
}

func reverse_aux(x uint64) uint64 {
	var t, q uint64
	var r uint64
	r = 0
	t = x
	for t >= 10 {
		q = t / 10
		r = r*10 + t-q*10
		t = q
	}
	r = r*10 + t
	return r
}
