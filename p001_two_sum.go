package goleetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, v := range nums {
		v1 := target - v
		if idx1,exist := m[v1]; exist {
			return []int{idx1, idx}
		}
		m[v] = idx
	}
	return nil
}