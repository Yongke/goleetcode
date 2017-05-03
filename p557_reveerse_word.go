package goleetcode

func reverseWords(s string) string {
	bs := []byte(s)
	length := len(s)
	ws := 0
	for i := 0; i < length; i++ {
		if bs[i] == ' ' {
			reverse_word_aux(bs[ws:i], i-ws)
			ws = i + 1
		}
	}
	reverse_word_aux(bs[ws:length], length-ws)
	return string(bs)
}

func reverse_word_aux(bs []byte, length int) {
	for i, j := length-1, 0; j < i; i, j = i-1, j+1 {
		bs[j], bs[i] = bs[i], bs[j]
	}
}
