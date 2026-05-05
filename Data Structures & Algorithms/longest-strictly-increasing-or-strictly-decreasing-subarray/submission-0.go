func longestMonotonicSubarray(nums []int) int {
	res := 1
	cnt := 1
	prev := nums[0]
	increasing := false
	decreasing := false
	for _, v := range nums[1:] {
		if prev < v {
			if increasing {
				cnt++
			} else {
				cnt = 2
			}
			increasing = true
			decreasing = false
		} else if prev > v {
			if decreasing {
				cnt++
			} else {
				cnt = 2
			}
			increasing = false
			decreasing = true
		} else {
			cnt = 1
			increasing = false
			decreasing = false
		}
		if cnt > res {
			res = cnt
		}
		prev = v
	}
	return res
}