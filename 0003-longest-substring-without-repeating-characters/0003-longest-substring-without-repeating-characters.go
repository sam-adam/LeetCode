func lengthOfLongestSubstring(s string) int {
	charCodes := make([]int, 256)
	startIndex := 0
	maxLen := 0

	for i := 0; i < len(s); i++ {
		charCode := s[i]
		biggestIndex := charCodes[charCode]

		if startIndex < biggestIndex {
			startIndex = biggestIndex
		}

		diff := i - startIndex + 1

		if diff > maxLen {
			maxLen = diff
		}

		charCodes[charCode] = i + 1
	}

	return maxLen
}
