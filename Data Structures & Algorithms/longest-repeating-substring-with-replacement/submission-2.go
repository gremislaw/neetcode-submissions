func characterReplacement(s string, k int) int {
	normalized := []rune(s)
	cnt := make(map[rune]int)
	res, maxCnt, l := 0, 0, 0
	for r := 0; r < len(normalized); r++ {
		cnt[normalized[r]]++
		if cnt[normalized[r]] > maxCnt {
			maxCnt = cnt[normalized[r]]
		}

		for r - l + 1 - maxCnt > k {
			cnt[normalized[l]]--
			l++
		}

		if r - l + 1 > res {
			res = r - l + 1
		}
	}
	return res
}
