func isValid(s string) bool {
	pairs := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	var openBrackets []string

	firstChar := string(s[0])
	if firstChar == ")" || firstChar == "}" || firstChar == "]" {
		return false
	}

	lastChar := string(s[len(s)-1])
	if lastChar == "(" || lastChar == "{" || lastChar == "[" {
		return false
	}

	for _, c := range s {
		cs := string(c)

		if cs == "(" || cs == "{" || cs == "[" {
			openBrackets = append(openBrackets, cs)
		} else {
			if len(openBrackets) == 0 {
				return false
			}

			lastOpenBracket := openBrackets[len(openBrackets)-1]
			supposedPair := pairs[lastOpenBracket]

			if cs != supposedPair {
				return false
			}

			if len(openBrackets) == 1 {
				openBrackets = []string{}
			} else {
				openBrackets = openBrackets[0:len(openBrackets)-1]
			}
		}
	}

	return len(openBrackets) == 0
}
