package goleetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	odd := 1 == (l1+l2)%2

	var last int
	for i, j, k := 0, 0, 0; i < l1 || j < l2; {
		var cur int
		if i >= l1 && j < l2 {
			cur = nums2[j]
			j++
			goto check
		}

		if j >= l2 && i < l1 {
			cur = nums1[i]
			i++
			goto check
		}

		if nums1[i] < nums2[j] {
			cur = nums1[i]
			i++
		} else {
			cur = nums2[j]
			j++
		}

	check:
		k++
		if k == (l1+l2)/2+1 {
			if odd {
				return float64(cur)
			}
			return (float64(cur) + float64(last)) / 2
		}
		last = cur
	}
	return 0
}
