func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

	chars := make([]string, (2*len(s))+1)

	chars[0] = "#"
	chars[len(chars)-1] = "#"

	for i, c := range s {
		chars[2*i] = string(c)

		if i < len(s)-1 {
			chars[2*i+1] = "#"
		}
	}

	s1 := strings.Join(chars, "")

	center := 0
	radius := 0
	radiuses := make([]int, len(s1))

	for center < len(s1) {
		leftBound := center - (radius + 1)
		rightBound := center + (radius + 1)

		for leftBound >= 0 && rightBound < len(s1) && s1[leftBound] == s1[rightBound] {
			radius++

			leftBound = center - (radius + 1)
			rightBound = center + (radius + 1)
		}

		radiuses[center] = radius

		oldCenter := center
		oldRadius := radius

		center++
		radius = 0

		for center <= oldCenter + oldRadius {
			mirroredCenter := oldCenter - (center - oldCenter)
			maxMirroredRadius := oldCenter + oldRadius - center

			if radiuses[mirroredCenter] < maxMirroredRadius {
				radiuses[center] = radiuses[mirroredCenter]
				center++
			} else if radiuses[mirroredCenter] > maxMirroredRadius {
				radiuses[center] = maxMirroredRadius
				center++
			} else {
				radius = maxMirroredRadius
				break
			}
		}
	}

	maxC := 0

	for c, r := range radiuses {
		if r > radiuses[maxC] {
			maxC = c
		}
	}

	palin := s1[maxC - radiuses[maxC]:maxC + radiuses[maxC] + 1]

	return strings.Join(strings.Split(palin, "#"), "")
}
