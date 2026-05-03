func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	stack := make([]int, 0, n)

	for i, temp := range temperatures {
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[prev] = i - prev
		}
		stack = append(stack, i)
	}
	return res
}