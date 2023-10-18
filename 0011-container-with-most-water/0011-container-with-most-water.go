func maxArea(height []int) int {
	maxArea := 0
	heightLen := len(height)
	i := 0
	j := heightLen - 1

	for i < j {
		area := calculateArea(height, i, j)

		if area > maxArea {
			maxArea = area
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	// for i := 0; i < heightLen - 1; i++ {
	// 	for j := i + 1; j < heightLen; j++ {
	//         if height[j] > height[i] && height[j-1] < height[i] && height[j-1] - height[i] > j-i {
	// 			i = j - 1
	// 		}

	// 		iHeight := height[i]
	// 		jHeight := height[j]
	// 		shortest := iHeight

	// 		if jHeight < iHeight {
	// 			shortest = jHeight
	// 		}

	// 		area := (j - i) * shortest

	//         // fmt.Println(fmt.Sprintf("i: %d, j: %d, iHeight: %d, jHeight: %d, area: %d, maxArea: %d", i, j, iHeight, jHeight, area, maxArea))

	// 		if area > maxArea {
	// 			maxArea = area
	// 		}
	// 	}
	// }

	return maxArea
}

func calculateArea(height []int, i int, j int) int {
	if height[i] < height[j] {
		return (j - i) * height[i]
	} else {
		{
			return (j - i) * height[j]
		}
	}
}
