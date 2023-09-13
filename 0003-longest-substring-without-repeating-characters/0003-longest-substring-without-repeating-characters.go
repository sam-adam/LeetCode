func lengthOfLongestSubstring(s string) int {
	var longest int

	arr := strings.Split(s, "")
	arrLen := len(arr)

	for i := 0; i < arrLen - longest; i++ {
		chars := make(map[string]int, 0)
		length := 1

		chars[arr[i]] = i

		for j := i + 1; j < arrLen; j++ {
			if dupIdx, exists := chars[arr[j]]; exists {
				i = dupIdx

				break
			}

			chars[arr[j]] = j
			length++
		}

		if length > longest {
			longest = length
		}
	}

	return longest
}
