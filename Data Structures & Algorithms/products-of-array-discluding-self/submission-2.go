func productExceptSelf(nums []int) []int {
	fullProduct := 1
	haveZero := 0
	for _, v := range nums {
		if v == 0 {
			haveZero++
			continue
		}
		fullProduct *= v
	}
	if haveZero > 1 {
		fullProduct = 0
	}
	res := make([]int, len(nums))
	for i, v := range nums {
		if v == 0 {
			res[i] = fullProduct
			continue
		}
		if haveZero > 0 {
			res[i] = 0
			continue
		}
		res[i] = fullProduct / v
	}
	return res
}
