func convert(s string, numRows int) string {
	gap := (2 * numRows) - 2

	if gap < 0 {
		gap = 0
	}

	var res []string

	for i := 0; i < numRows; i++ {
		step := gap - (i * 2)

		if step == 0 {
			step = gap
		}

		for j := i; j < len(s); {
			res = append(res, string(s[j]))

			if step > 0 {
				j += step
			} else {
				j++
			}

			step = gap - step

			if step == 0 {
				step = gap
			}
		}
	}

	return strings.Join(res, "")
}
