func merge(nums1 []int, m int, nums2 []int, n int) {
	res := make([]int, len(nums1))
	i1, i2 := 0, 0
	for i := 0; i < len(res); i++ {
		if i1 == m {
			res[i] = nums2[i2]
			i2++
			continue
		}
		if i2 == n {
			res[i] = nums1[i1]
			i1++
			continue
		}
		if nums1[i1] < nums2[i2] {
			res[i] = nums1[i1]
			i1++
		} else {
			res[i] = nums2[i2]
			i2++
		}
	}
	for i := 0; i < len(res); i++ {
		nums1[i] = res[i]
	}
}
