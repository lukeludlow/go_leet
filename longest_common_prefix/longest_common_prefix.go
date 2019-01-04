package goleet

func longestCommonPrefix(strs []string) string {
	prefix := ""
	if len(strs) == 0 {
		return prefix
	}
	length := len(strs[0])
	for _, s := range strs {
		length = min(length, len(s))
	}
	for i := 0; i < length; i++ {
		char := strs[0][i]
		for _, s := range strs {
			if s[i] != char {
				return prefix
			}
		}
		prefix += string(char)
	}
	return prefix
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
