package goleetcode

func reverseStr(s string, k int) string {
	bs := []byte(s)
	length := len(s)

	for i := 0; i < length; i = i + 2*k {
		if i+k > length {
			tbs := bs[i:]
			reverse_str_aux(tbs, length-i)
		} else {
			tbs := bs[i : i+k]
			reverse_str_aux(tbs, k)
		}
	}
	return string(bs)
}

func reverse_str_aux(bs []byte, length int) {
	for i, j := length-1, 0; j < i; i, j = i-1, j+1 {
		bs[j], bs[i] = bs[i], bs[j]
	}
}
