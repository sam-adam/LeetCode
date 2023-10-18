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
