func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    check := make([]int, 26)
    
    for i := 0; i < len(s); i++ {
        check[s[i]-'a']++
        check[t[i]-'a']--
    }

    for _, c := range check {
        if c != 0 {
            return false
        }
    }

    return true
}
