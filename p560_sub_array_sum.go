package goleetcode

import "fmt"

func subarraySum(nums []int, k int) int {
	l := len(nums)
	sumi := 0

	tmp := make(map[int][]int)
	for i := 0; i < l; i++ {
		sumi += nums[i]
		tmp[sumi] = append(tmp[sumi], i)
	}

	match := 0
	fmt.Println(tmp)
	for s, idx := range tmp {
		if k == s {
			match = match + len(idx)
		}
		if idx1, exist := tmp[s+k]; exist {
			match = match + compare_idx(idx, idx1)
		}
	}
	return match
}

func compare_idx(idx_small []int, idx_big []int) int {
	x := 0
	for _, b := range idx_big {
		for _, s := range idx_small {
			if b > s {
				x++
			}
		}
	}
	return x
}
