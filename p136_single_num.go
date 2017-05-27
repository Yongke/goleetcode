package goleetcode

func singleNumber(nums []int) int {
	i := 0
	for _, v := range nums {
		i ^= v
	}
	return i
}
