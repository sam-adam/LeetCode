func reverse(x int) int {
	// The result variable
	var reversed int

	// Check for negative value, x will be converted to positive value if
	// it's negative, but will be converted back to negative in the end
	negativeMultiplier := 1

	if x < 0 {
		negativeMultiplier = -1
	}

	x = x * negativeMultiplier

	minPow := 0
	maxPow := 0

	// Get x total numbers by checking against 10^i, similar to strlen
	for i := 0; int(math.Pow10(i)) <= x; i++ {
		maxPow = i
	}

	// Starts from the biggest original digit until the lowest digit
	for i := maxPow; i >= minPow; i-- {
		// Get the digit in current position, e.g:
		// For x = 567 and i = 2, num = 5
		// For x = 567 and i = 1, num = 6
		// For x = 567 and i = 0, num = 7
		num := x % int(math.Pow10(i + 1)) / int(math.Pow10(i))

		// Increment result var by multiplying num to the lower mirror digit, e.g:
		// For x = 567 and i = 2 and num = 5, result = 5 * 10^0
		// For x = 567 and i = 1 and num = 6, result = 6 * 10^1
		// For x = 567 and i = 0 and num = 7, result = 7 * 10^2
		reversed += num * int(math.Pow10(maxPow - i))
	}

	// Revert the result's negative value if it was negative
	reversed = reversed * negativeMultiplier

	// Final check to make sure the result is within int32
	if reversed >= math.MinInt32 && reversed <= math.MaxInt32 {
		return reversed
	} else {
		return 0
	}
}
