func maxArea(heights []int) int {
	l, r := 0, len(heights) - 1
	maxV := 0
	for ; l < r; {
		v1 := min(heights[l], heights[r]) * (r-l)
		if v1 > maxV {maxV = v1}
		if heights[l] > heights[r] {
			r--
		} else {
			l++
		}
	}
	return maxV
}
