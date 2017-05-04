package goleetcode

// first version: o(n^2) ~~ n(n+1)/2
func subarraySum(nums []int, k int) int {
	l := len(nums)
	match := 0
	tmp := make([]int, l)
	for i := 0; i < l; i++ {
		for j := 0; j <= i; j++ {
			tmp[j] += nums[i]
			if tmp[j] == k {
				match++
			}
		}
	}
	return match
}
