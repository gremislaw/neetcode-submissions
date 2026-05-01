type Frequent struct {
	num int
	cnt int
}

func topKFrequent(nums []int, k int) []int {
	seen := make([]Frequent, 2005, 2005)
	for _, v := range nums {
		seen[v + 1000].num = v
		seen[v + 1000].cnt++
	}
	sort.Slice(seen, func(i, j int) bool {
		return seen[i].cnt > seen[j].cnt
	})
	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, seen[i].num)
	}
	return res
}
