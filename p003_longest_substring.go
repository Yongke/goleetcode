package goleetcode

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	maxsublen := 0
	sublen := 0

	for i := 0; i < len(s); i++ {
		val := s[i]
		if last_i, exist := m[val]; exist {
			if sublen > maxsublen {
				maxsublen = sublen
			}

			m = make(map[byte]int)
			for j := last_i + 1; j <= i; j++ {
				m[s[j]] = j
			}
			sublen = i - last_i
			continue
		}

		m[val] = i
		sublen += 1
	}
	if sublen > maxsublen {
		maxsublen = sublen
	}
	return maxsublen
}