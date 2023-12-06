func generateParenthesis(n int) []string {
	var result []string
	var backtrack func(s string, left int, right int)

	backtrack = func(s string, left int, right int) {
		if len(s) == 2*n {
			result = append(result, s)
			return
		}
		if left < n {
			backtrack(s+"(", left+1, right)
		}
		if right < left {
			backtrack(s+")", left, right+1)
		}
	}

	backtrack("", 0, 0)

	return result
}