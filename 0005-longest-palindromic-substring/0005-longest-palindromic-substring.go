func longestPalindrome(s string) string {
	longest := ""
	longestLen := 0

	for leftIdx := range s {
		for rightIdx := len(s) - 1; rightIdx >= leftIdx + longestLen; rightIdx-- {
			if !checkIsPalindrome(s, leftIdx, rightIdx) {
				continue
			}

			palinLen := rightIdx + 1 - leftIdx

			if palinLen > longestLen {
				longest = s[leftIdx:rightIdx+1]
				longestLen = palinLen
			}

			break
		}
	}

	return longest
}

func checkIsPalindrome(s string, leftIdx int, rightIdx int) bool {
	isMirrorMatch := s[leftIdx] == s[rightIdx]

	for rightIdx > leftIdx {
		if !isMirrorMatch {
			return false
		}

		leftIdx += 1
		rightIdx -= 1

		isMirrorMatch = s[leftIdx] == s[rightIdx]
	}

	return true
}
