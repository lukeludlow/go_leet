package goleet

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	xStr := strconv.Itoa(x)
	length := len(xStr)
	for i := 0; i < length; i++ {
		if xStr[i] != xStr[length-1-i] {
			return false
		}
	}
	return true
}
