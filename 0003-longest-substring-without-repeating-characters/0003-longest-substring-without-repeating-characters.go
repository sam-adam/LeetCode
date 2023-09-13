func lengthOfLongestSubstring(s string) int {
    dict := make([]int, 256)
    maxLen := 0
    start := 0
    
    for i := 0; i < len(s); i++ {
        if start < dict[s[i]] {
            start = dict[s[i]]
        }
        if maxLen < i - start + 1 {
            maxLen = i - start + 1
        }
        dict[s[i]] = i + 1
    }
    return maxLen
}