func romanToInt(s string) int {
	result := 0
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	symbolsMap := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}

	for _, symbol := range symbols {
		for len(s) >= len(symbol) && s[0:len(symbol)] == symbol {
			result += symbolsMap[symbol]

			if len(s) > 1 {
				s = s[len(symbol):]
			} else {
				break
			}
		}
	}

	return result
}
