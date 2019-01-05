package goleet

func isValid(s string) bool {
	var stack []rune
	for _, c := range s {
		switch c {
		case ')':
			if len(stack) == 0 {
				return false
			}
			// pop
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != '(' {
				return false
			}
		case ']':
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != '[' {
				return false
			}
		case '}':
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != '{' {
				return false
			}
		default:
			// push
			stack = append(stack, c)
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
