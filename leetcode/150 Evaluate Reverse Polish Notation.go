func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		switch token {
		case "+":
			n1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			n2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, n2+n1)
		case "-":
			n1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			n2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, n2-n1)
		case "*":
			n1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			n2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, n2*n1)
		case "/":
			n1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			n2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, n2/n1)
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[0]
}