package goleetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	s1, s2 := len(nums1), len(nums2)
	if (s1+s2)&1 == 1 {
		return float64(findKth(nums1, nums2, (s1+s2)/2+1))
	}
	return float64(findKth(nums1, nums2, (s1+s2)/2)+findKth(nums1, nums2, (s1+s2)/2+1)) / 2
}

func findKth(nums1 []int, nums2 []int, k int) int {
	s1, s2 := len(nums1), len(nums2)

	if s1 == 0 {
		return nums2[k-1]
	}
	if s2 == 0 {
		return nums1[k-1]
	}
	if k == 1 {
		return min(nums1[0], nums2[0])
	}
	max1, n1 := maxKth(nums1, k/2)
	max2, n2 := maxKth(nums2, k/2)

	if max1 < max2 {
		return findKth(nums1[n1:], nums2, k-n1)
	}
	return findKth(nums1, nums2[n2:], k-n2)
}

func maxKth(nums []int, k int) (int, int) {
	if k > len(nums) {
		return nums[len(nums)-1], len(nums)
	}
	return nums[k-1], k
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
