package goleetcode

func reverseString(s string) string {
	length := len(s)
	bs := []byte(s)
	for i, j := length-1, 0; j < i; i, j = i-1, j+1 {
		bs[j],bs[i] = s[i],s[j]
	}
	return string(bs)
}
