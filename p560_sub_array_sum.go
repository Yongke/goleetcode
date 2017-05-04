package goleetcode

func subarraySum(nums []int, k int) int {
	l := len(nums)
	sumi, match := 0, 0
	m := make(map[int]int)
	m[0] = 1 //what the fuck!
	for i := 0; i < l; i++ {
		sumi += nums[i]
		pre_sumi := sumi - k
		if cnt, exist := m[pre_sumi]; exist {
			match += cnt
		}
		m[sumi]++
	}
	return match
}
