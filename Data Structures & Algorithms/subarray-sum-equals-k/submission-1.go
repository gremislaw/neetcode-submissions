func subarraySum(nums []int, k int) int {
	res := 0
	prefSumMap := make(map[int]int)
	prefSumMap[0] = 1
	prefSum := 0
	for _, v := range nums {
		prefSum += v
		if cnt, ok := prefSumMap[prefSum-k]; ok {
			res += cnt
		}
		prefSumMap[prefSum]++
	}
	return res
}
