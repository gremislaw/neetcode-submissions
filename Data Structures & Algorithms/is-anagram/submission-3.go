func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	seen := make(map[rune]int)
	for _, val := range s {
		seen[val]++
	}
	for _, val := range t {
		if r, exists := seen[val]; exists && r > 0 {
			seen[val]--
			continue
		}
		return false
	}
	return true
}
