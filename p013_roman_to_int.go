package goleetcode

func romanToInt(s string) int {
	m := map[string]int{"M": 1000, "MM": 2000, "MMM": 3000,
		"C": 100, "CC": 200, "CCC": 300, "CD": 400, "D": 500, "DC": 600, "DCC": 700, "DCCC": 800, "CM": 900,
		"X": 10, "XX": 20, "XXX": 30, "XL": 40, "L": 50, "LX": 60, "LXX": 70, "LXXX": 80, "XC": 90,
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9}
	buff := []byte{}
	last, sum := 0, 0
	for i := 0; i < len(s); i++ {
		buff = append(buff, s[i])
		val, exist := m[string(buff)]
		if exist {
			last = val
		} else {
			sum += last
			buff = []byte{s[i]}
			last = m[string(buff)]
		}
	}
	return sum + last
}
