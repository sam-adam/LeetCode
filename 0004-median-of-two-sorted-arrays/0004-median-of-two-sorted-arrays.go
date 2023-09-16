
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	isOdd := totalLen%2 == 1 // we will separate odd / even solution for execution time's sake
	lenNums1 := len(nums1) - 1
	lenNums2 := len(nums2) - 1

  // first we start both pointers to -1, which we will increment
  // by 1 for each step, depends on which is lower and have next
  // element
	p1 := -1
	p2 := -1

	if isOdd {
    // for odd m+n, we will traverse just until ((m+n)/2) - 1
		medianIdx := int(math.Floor(float64(totalLen) / 2.0))

		for i := 0; i < medianIdx; i++ {
			if p1+1 > lenNums1 {
				p2++
			} else if p2+1 > lenNums2 {
				p1++
			} else if nums1[p1+1] <= nums2[p2+1] {
				p1++
			} else {
				p2++
			}
		}

		if p1+1 > lenNums1 {
			return float64(nums2[p2+1])
		} else if p2+1 > lenNums2 {
			return float64(nums1[p1+1])
		} else if nums1[p1+1] < nums2[p2+1] {
			return float64(nums1[p1+1])
		} else {
			return float64(nums2[p2+1])
		}
	} else {
		medianIdx := totalLen / 2
		numPre := 0

		for i := 0; i < medianIdx; i++ {
			if p1+1 > lenNums1 {
				p2++
				numPre = nums2[p2]
			} else if p2+1 > lenNums2 {
				p1++
				numPre = nums1[p1]
			} else if nums1[p1+1] <= nums2[p2+1] {
				p1++
				numPre = nums1[p1]
			} else {
				p2++
				numPre = nums2[p2]
			}
		}

		if p1+1 > lenNums1 {
			return (float64(nums2[p2+1]) + float64(numPre)) / 2.0
		} else if p2+1 > lenNums2 {
			return (float64(nums1[p1+1]) + float64(numPre)) / 2.0
		} else if nums1[p1+1] < nums2[p2+1] {
			return (float64(nums1[p1+1]) + float64(numPre)) / 2.0
		} else {
			return (float64(nums2[p2+1]) + float64(numPre)) / 2.0
		}
	}
}