const zag = 2

func convert(s string, numRows int) string {
    // Calculate the gap given the numRows using the function
    // 
    // y = max(0, nx - n)
    // 
    // Where x is the numRows, n is the zag constant, and y is the resulting gap.
    //
    // Examples:
    // NumRows = 1, Gap = 0
    // NumRows = 2, Gap = 2
    // NumRows = 3, Gap = 4
    // NumRows = 3, Gap = 6
    // NumRows = 4, Gap = 8
    // NumRows = 5, Gap = 10
    // etc..
	gap := (zag * numRows) - zag

    // Guard the gap so it has a min value of 0.
    // Go's Math.max function requires and outputs float, 
    // too troublesome to convert before and after.
	if gap < 0 {
		gap = 0
	}

	var res []string

    // The outer loop, will loop from 0..numRows - 1
	for i := 0; i < numRows; i++ {
        // Calculate how much the pointer needs to be shifted on the next step.
		step := gap - (i * zag)

        // If step is 0, meaning we've reached the bottom-most row, which step is equal to the gap variable.
        // For example, given numRows = 5, then the steps for each row would be:
        // i = 0, 1st step = 8 - (0 * 2) = 8
        // i = 0, 2nd step = 8 - 1st step = 8
        // i = 0, 3rd step = 8 - 2nd step = 8
        // i = 0, etc..
        // i = 1, 1st step = 8 - (1 * 2) = 6
        // i = 1, 2nd step = 8 - 1st step = 2
        // i = 1, 3rd step = 8 - 2nd step = 6
        // i = 1, etc..
		if step == 0 {
			step = gap
		}

		for j := i; j < len(s); {
			res = append(res, string(s[j]))

            // This ensure that j will always be incremented by at least 1, to guard cases where numRows = 1
			if step > 0 {
				j += step
			} else {
				j++
			}

            // Switch the step variable between going down the valley, or climbing up the peak
			step = gap - step

			if step == 0 {
				step = gap
			}
		}
	}

	return strings.Join(res, "")
}
