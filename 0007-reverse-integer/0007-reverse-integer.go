func reverse(x int) int {
	var reversed int

	isNegative := x < 0
	negativeMultiplier := 1

	if isNegative {
		negativeMultiplier = -1
	}

	x = x * negativeMultiplier

	minPow := 0
	maxPow := 0

	for i := 0; int(math.Pow10(i)) <= x; i++ {
		maxPow = i
	}

	for i := maxPow; i >= minPow; i-- {
		num := x % int(math.Pow10(i + 1)) / int(math.Pow10(i))

		reversed += num * int(math.Pow10(maxPow - i))
	}

	reversed = reversed * negativeMultiplier

	if reversed >= math.MinInt32 && reversed <= math.MaxInt32 {
		return reversed
	} else {
		return 0
	}
}
