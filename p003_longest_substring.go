package goleetcode

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l == 0 || l == 1 {
		return l
	}

	m := [128]int{}
	res, tmp := 0, 0
	last_idx := 0
	for i := 0; i < l; i++ {
		c := s[i]
		idx := max(m[c], last_idx)
		if idx != 0 {
			tmp = i + 1 - idx
			last_idx = idx
		} else {
			tmp++
		}

		res = max(res, tmp)
		m[c] = i + 1
	}
	return res
}

func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}
