func lengthOfLongestSubstring(s string) int {
	set := make(map[rune]struct{})
	for _, v := range s {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
		}
	}
	windowLen := len(set)
	res := 0
	normalized := []rune(s)
	for i := 0; i < len(normalized); i++ {
		newSet := make(map[rune]struct{})
		for j := i; j < min(i + windowLen, len(normalized)); j++ {
			if _, ok := newSet[normalized[j]]; !ok {
				newSet[normalized[j]] = struct{}{}
			} else {
				break
			}
		}
		if len(newSet) > res {
			res = len(newSet)
		}
	}
	return res
}
