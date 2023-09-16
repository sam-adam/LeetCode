func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	isOdd := totalLen%2 == 1
	lenNums1 := len(nums1) - 1
	lenNums2 := len(nums2) - 1

	p1 := -1
	p2 := -1

	fmt.Println(fmt.Sprintf("nums1: %v", nums1))
	fmt.Println(fmt.Sprintf("nums2: %v", nums2))

	if isOdd {
		fmt.Println("odd")
		medianIdx := int(math.Floor(float64(totalLen) / 2.0))

		for i := 0; i < medianIdx; i++ {
			if p1+1 > lenNums1 {
				p2++
			} else if p2+1 > lenNums2 {
				p1++
			} else if nums1[p1+1] <= nums2[p2+1] {
				fmt.Println(fmt.Sprintf("p1: %d -> %d", p1, p1+1))
				p1++
			} else {
				fmt.Println(fmt.Sprintf("p2: %d -> %d", p2, p2+1))
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
		fmt.Println("even")
		medianIdx := totalLen / 2
		numPre := 0

		fmt.Println(fmt.Sprintf("median-idx: %d", medianIdx))

		for i := 0; i < medianIdx; i++ {
			fmt.Println(fmt.Sprintf("i: %d, p1+1: %d, lenNums1: %d, p2+1: %d, lenNums2: %d", i, p1+1, lenNums1, p2+1, lenNums2))

			if p1+1 > lenNums1 {
				p2++
				numPre = nums2[p2]
			} else if p2+1 > lenNums2 {
				p1++
				numPre = nums1[p1]
			} else if nums1[p1+1] <= nums2[p2+1] {
				fmt.Println(fmt.Sprintf("p1: %d -> %d", p1, p1+1))
				p1++
				numPre = nums1[p1]
			} else {
				fmt.Println(fmt.Sprintf("p2: %d -> %d", p2, p2+1))
				p2++
				numPre = nums2[p2]
			}
		}

		if p1+1 > lenNums1 {
			return float64((float64(nums2[p2+1]) + float64(numPre)) / 2)
		} else if p2+1 > lenNums2 {
			return float64((float64(nums1[p1+1]) + float64(numPre)) / 2)
		} else if nums1[p1+1] < nums2[p2+1] {
			return float64((float64(nums1[p1+1]) + float64(numPre)) / 2)
		} else {
			return float64((float64(nums2[p2+1]) + float64(numPre)) / 2)
		}
	}
}
