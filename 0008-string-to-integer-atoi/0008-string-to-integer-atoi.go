func myAtoi(s string) int {
	result := uint64(0)
	pow := 0
	isOverflow := false
	negativeMultiplier := 1
	nums := map[string]uint64{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}

	for i := len(s) - 1; i >= 0; i-- {
		str := string(s[i])
		n, isNum := nums[str]

		if !isNum {
			shouldReset := (str != " " && str != "-" && str != "+") ||
				(i < len(s)-1 && (str == "-" || str == "+") && (string(s[i+1]) == "+" || string(s[i+1]) == "-" || string(s[i+1]) == " "))

			if shouldReset {
				result = 0
				pow = 0
				isOverflow = false
			} else if str == "-" {
				negativeMultiplier = -1
			} else if str != " " {
				negativeMultiplier = 1
			}
		} else {
			shouldReset := i < len(s)-1 && (string(s[i+1]) == "+" || string(s[i+1]) == "-" || string(s[i+1]) == " ")

			if shouldReset {
				result = 0
				pow = 0
				isOverflow = false
			}

			negativeMultiplier = 1
			result += n*uint64(math.Pow10(pow))
			isOverflow = isOverflow || result > math.MaxInt32
			pow++
		}
	}

	if negativeMultiplier == -1 && result > math.MaxInt32 {
		return math.MinInt32
	}

	if isOverflow {
		return math.MaxInt32
	}

	if result > math.MaxInt32 {
		if negativeMultiplier == -1 {
			return math.MinInt32 - 1
		} else {
			return math.MaxInt32
		}
	}

	return int(result) * negativeMultiplier
}
