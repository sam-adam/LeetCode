func isPalindrome(x int) bool {
	strX := strconv.Itoa(x)
	chars := make([]string, (2*len(strX))+1)

	chars[len(chars)-1] = "#"

	for _, c := range strX {
		chars = append(chars, string(c), "#")
	}

	s1 := strings.Join(chars, "")
	center := len(s1) / 2
	left := center
	right := center

	for left > 0 {
		if string(s1[left]) != string(s1[right]) {
			return false
		}

		left--
		right++
	}

	return true
}
