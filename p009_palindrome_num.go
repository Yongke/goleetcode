package goleetcode

func isPalindrome(x int) bool {
	if x >= 0 && x < 10 {
		return true
	}
	if x < 0 || (x%10 == 0) {
		return false
	}

	var t, q, r = 0, 0, 0
	for x >= 10 {
		q = x / 10
		t = x - q*10
		r = r*10 + t
		if r == q || r == x {
			return true
		}
		x = q
	}
	return false
}
