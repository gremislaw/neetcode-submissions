func isPalindrome(s string) bool {
	var res strings.Builder
	res.Grow(len(s))
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			res.WriteRune(r)
		}
	}
	tmp := strings.ToLower(res.String())
	fmt.Println(tmp)
	for i, j := 0, len(tmp) - 1; i < j; i, j = i + 1, j - 1 {
		if string(tmp[i]) != string(tmp[j]) {
			return false
		}
	}
	return true
}
