func lengthOfLongestSubstring(s string) int {
	var longest int

	arr := strings.Split(s, "")
	arrLen := len(arr)

	for i := 0; i < arrLen; i++ {
		chars := make(map[string]bool, 0)
		length := 1

		chars[arr[i]] = true

		for j := i + 1; j < arrLen; j++ {
			if _, exists := chars[arr[j]]; exists {
				break
			}

			chars[arr[j]] = true
			length++
		}

		if length > longest {
					longest = length
				}
	}

	return longest
}
