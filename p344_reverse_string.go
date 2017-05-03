package goleetcode

func reverseString(s string) string {
	l := len(s)
	bs := make([]byte,l)
	for i := l-1; i >= 0; i-- {
		bs[l-1-i] = s[i]
	}
	return string(bs)
}
