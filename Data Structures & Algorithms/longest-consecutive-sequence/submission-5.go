import "slices"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slices.Sort(nums)
	prevV := nums[0]
	cnt := 1
	maxCnt := 1
	for _, v := range nums[1:] {
		if prevV+1 == v {
			cnt++
			if cnt > maxCnt {
				maxCnt = cnt
			}
		} else if prevV == v {
			continue
		} else {
			cnt = 1
		}
		prevV = v
	}
	return maxCnt
}
