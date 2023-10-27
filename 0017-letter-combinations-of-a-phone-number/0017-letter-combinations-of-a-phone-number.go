func letterCombinations(digits string) []string {
	keyboard := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}

	if digits == "" {
		return []string{}
	}

	if len(digits) == 1 {
		return keyboard[digits]
	}

    words := keyboard[string(digits[0])]

	for i := 1; i < len(digits); i++ {
        var newWords []string

        chars := keyboard[string(digits[i])]

        for _, c := range chars {
            for _, w := range words {
                newWords = append(newWords, w + c)
            }
        }

        words = newWords
    }

	return words
}