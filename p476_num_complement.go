package goleetcode

func findComplement(num int) int {
	if num <= 0 {
		panic("invalid input")
	}
	var i uint32
	for i = 0; i < 32; i++ {
		if num>>i == 1 {
			return num ^ (1<<(i+1) - 1)
		}
	}
	// not possible here
	return 0
}
