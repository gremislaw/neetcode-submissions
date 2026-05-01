func twoSum(nums []int, target int) []int {
    set := make(map[int]int)
	for i, num := range nums {
		set[num] = i
	}
	for i, num := range nums {
		if j, exists := set[target - num]; exists && i != j {
			if i < j { return []int{i, j} }
			return []int{j, i}
		}
	}
	return []int{}
}
