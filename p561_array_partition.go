package goleetcode

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	var sum = 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}
