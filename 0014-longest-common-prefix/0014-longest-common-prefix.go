func longestCommonPrefix(strs []string) string {
    longest := ""
    shortestLen := len(strs[0])

    if len(strs) == 1 {
        return strs[0]
    }

    for j := 1; j < len(strs); j++ {
        hasSamePrefix := false

        if len(strs[j]) < shortestLen {
            shortestLen = len(strs[j])
        }

        for k := shortestLen; k > 0 && hasSamePrefix == false; k-- {
            shortestLen = k
            strI := string(strs[0][0:k])
            strJ := string(strs[j][0:k])

            if strI == strJ {
                longest = strI
                hasSamePrefix = true

                break
            }
        }

        if hasSamePrefix == false {
            return ""
        }
    }

    return longest
}