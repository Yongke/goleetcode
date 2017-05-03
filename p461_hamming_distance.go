package goleetcode

func hammingDistance(x int, y int) int {
	d := 0
	var i uint8
	for i = 0; i <= 31; i++ {
		mask := 1 << i
		if (x & mask) != (y & mask) {
			d++
		}
	}
	return d
}
