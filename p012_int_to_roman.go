package goleetcode

import (
	"strings"
)

var m = map[int]string{
	1: "I",
	2: "II",
	3: "III",
	4: "IV",
	5: "V",
	6: "VI",
	7: "VII",
	8: "VIII",
	9: "IX",
}

func intToRoman(num int) string {
	if num >= 4000 {
		panic("not support")
	}
	buff := intToRomanAux(num)
	s := len(buff)
	for i := 0; i < s/2; i++ {
		buff[i], buff[s-i-1] = buff[s-i-1], buff[i]
	}
	return strings.Join(buff, "")
}

func intToRomanAux(num int) []string {
	if num >= 1000 {
		t := num / 1000
		return append(intToRomanAux(num-t*1000), strings.Repeat("M", t))
	} else if num >= 900 && num < 1000 {
		return append(intToRomanAux(num-900), "CM")

	} else if num >= 500 && num < 900 {
		return append(intToRomanAux(num-500), "D")

	} else if num >= 400 && num < 500 {
		return append(intToRomanAux(num-400), "CD")

	} else if num >= 100 && num < 400 {
		t := num / 100
		return append(intToRomanAux(num-t*100), strings.Repeat("C", t))

	} else if num >= 90 && num < 100 {
		return append(intToRomanAux(num-90), "XC")

	} else if num >= 50 && num < 90 {
		return append(intToRomanAux(num-50), "L")

	} else if num >= 40 && num < 50 {
		return append(intToRomanAux(num-40), "XL")

	} else if num >= 10 && num < 40 {
		t := num / 10
		return append(intToRomanAux(num-t*10), strings.Repeat("X", t))

	}
	return []string{m[num]}
}
