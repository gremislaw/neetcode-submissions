import "slices"

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)

	slices.Sort(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			sum := nums[i]+nums[j]+nums[k]
			if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for ; j < k && nums[j] == nums[j+1]; j++ {}
				for ; j < k && nums[k] == nums[k-1]; k-- {}
				j++
				k--
			}
		}
	}
	return res
}
