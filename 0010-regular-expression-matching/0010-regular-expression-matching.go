func isMatch(s string, p string) bool {
	sRe := regexp.MustCompile(`(?m)\*{2,}`)
	p = sRe.ReplaceAllString(p, "*")
	re := regexp.MustCompile(`(?m)^` + p + `$`)
	matches := re.FindAllString(s, -1)

	return len(matches) > 0
}
