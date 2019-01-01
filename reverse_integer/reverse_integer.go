package goleet

func reverse(x int) int {
	ret := 0
	signed := 0
	if x < 0 {
		signed = -1
		x = -x
	} else {
		signed = 1
	}
	for x > 0 {
		ret = ret*10 + x%10
		x = x / 10
	}
	maxInt := 2147483647
	minInt := -2147483648
	if ret > maxInt || ret < minInt {
		return 0
	}
	return ret * signed
}
