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
	maxPalin := ""

	for center < len(s1) {
		radius := 0
		leftBound := center - (radius + 1)
		rightBound := center + (radius + 1)

		for leftBound >= 0 && rightBound < len(s1) && s1[leftBound] == s1[rightBound] {
			radius++

			palin := s1[leftBound:rightBound+1]

			if len(palin) > len(maxPalin) {
				maxPalin = palin
			}

			leftBound = center - (radius + 1)
			rightBound = center + (radius + 1)
		}

		center++
	}

	return strings.Join(strings.Split(maxPalin, "#"), "")
}
