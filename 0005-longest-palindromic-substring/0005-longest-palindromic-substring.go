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

		for center - (radius + 1) >= 0 && center + (radius + 1) < len(s1) && s1[center - (radius + 1)] == s1[center + (radius + 1)] {
			if len(s1[center - (radius + 1):center + (radius + 1)+1]) > len(maxPalin) {
				maxPalin = s1[center - (radius + 1):center + (radius + 1)+1]
			}

			radius++
		}

		center++
	}

	return strings.Join(strings.Split(maxPalin, "#"), "")
}
