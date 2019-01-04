package goleet

func romanToInt(s string) int {
	m := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	sum := 0
	var p rune // previous
	for i, c := range s {
		val := m[c]
		sum += val
		previousVal := m[p]
		p = c
		if i > 0 && val > previousVal {
			sum -= 2 * previousVal
		}
	}
	return sum
}
