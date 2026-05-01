import "slices"

func groupAnagrams(strs []string) [][]string {
	seen := make(map[string]int)
	anagrams := make([][]string, 0, 0)
	for _, str := range strs {
		tmp := []rune(str)
		slices.Sort(tmp)
		s := string(tmp)
		if i, ok := seen[s]; ok {
			anagrams[i] = append(anagrams[i], str)
		} else {
			nextInd := len(anagrams)
			anagrams = append(anagrams, []string{str})
			seen[s] = nextInd
		}
	}
	return anagrams
}
