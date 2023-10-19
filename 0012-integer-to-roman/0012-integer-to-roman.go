func intToRoman(num int) string {
	result := ""

	for num >= 1000 {
		result = result + "M"
		num -= 1000
	}

	for num >= 900 {
		result = result + "CM"
		num -= 900
	}

	for num >= 500 {
		result = result + "D"
		num -= 500
	}

	for num >= 400 {
		result = result + "CD"
		num -= 400
	}

	for num >= 100 {
		result = result + "C"
		num -= 100
	}

	for num >= 90 {
		result = result + "XC"
		num -= 90
	}

	for num >= 50 {
		result = result + "L"
		num -= 50
	}

	for num >= 40 {
		result = result + "XL"
		num -= 40
	}

	for num >= 10 {
		result = result + "X"
		num -= 10
	}

	for num >= 9 {
		result = result + "IX"
		num -= 10
	}

	for num >= 5 {
		result = result + "V"
		num -= 5
	}

	for num >= 4 {
		result = result + "IV"
		num -= 4
	}

	for num >= 1 {
		result = result + "I"
		num--
	}

	return result
}
