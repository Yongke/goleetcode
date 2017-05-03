package goleetcode

func matrixReshape(nums [][]int, r int, c int) [][]int {
	r1, c1 := len(nums), len(nums[0])
	if r1*c1 != r*c {
		return nums
	}

	row := make([]int, c)
	m := make([][]int, r)
	i, j := 0, 0
	for _, oldrow := range nums {
		for _, val := range oldrow {
			row[i] = val
			i++
			if i == c {
				i = 0
				m[j] = row
				j++
				row = make([]int, c)
			}
		}
	}
	return m
}
