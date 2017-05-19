package goleetcode

func nextGreaterElement(findNums []int, nums []int) []int {
	m := map[int]int{}
	m1 := map[int]int{}
	l := len(findNums)
	for i := 0; i < l; i++ {
		m[findNums[i]] = -1
	}

	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		for pre, _ := range m1 {
			if cur > pre {
				delete(m1, pre)
				m[pre] = cur
			}
		}
		if _, exist := m[cur]; exist {
			m1[cur] = cur
		}
	}
	result := make([]int, l)
	for i := 0; i < l; i++ {
		result[i] = m[findNums[i]]
	}
	return result
}
