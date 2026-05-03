func trap(height []int) int {
	l, r := 0, len(height)-1
	maxL, maxR := height[l], height[r]
	res := 0

	for l < r {
		if maxL < maxR {
			l++
			if height[l] < maxL {
				res += maxL - height[l]
			} else {
				maxL = height[l]
			}
		} else {
			r--
			if height[r] < maxR {
				res += maxR - height[r]
			} else {
				maxR = height[r]
			}
		}
	}
	return res
}