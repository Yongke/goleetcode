package goleetcode

import (
	"bytes"
)

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	maxsublen := 0
	sublen := 0
	bstmp := []byte{}

	for i := 0; i < len(s); i++ {
		val := s[i]
		if _, exist := m[val]; exist {
			if sublen > maxsublen {
				maxsublen = sublen
			}
			parts := bytes.Split(bstmp, []byte{val})
			for _, valOfPart1 := range parts[0] {
				delete(m, valOfPart1)
			}
			sublen = len(parts[1]) + 1
			bstmp = append(parts[1], val)
			continue
		}
		m[val] = 0
		sublen += 1
		bstmp = append(bstmp, val)
	}
	if sublen > maxsublen {
		maxsublen = sublen
	}
	return maxsublen
}