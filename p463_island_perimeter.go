package goleetcode

func islandPerimeter(grid [][]int) int {
	var last_row []int
	last_row = nil

	result := 0
	for _, row := range grid {
		last_flag := 0
		for i, flag := range row {
			if flag == 1 {
				result += 4
				if last_row != nil && last_row[i] == 1 {
					result -= 2
				}
				if last_flag == 1 {
					result -= 2
				}
			}
			last_flag = flag
		}
		last_row = row
	}
	return result
}
