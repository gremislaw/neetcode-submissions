func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens)/2+1)
	
	for _, t := range tokens {
		switch t {
		case "+", "-", "*", "/":
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-1]
			switch t {
			case "+": stack[len(stack)-1] = a + b
			case "-": stack[len(stack)-1] = a - b
			case "*": stack[len(stack)-1] = a * b
			case "/": stack[len(stack)-1] = a / b
			}
		default:
			v, _ := strconv.Atoi(t)
			stack = append(stack, v)
		}
	}
	return stack[0]
}