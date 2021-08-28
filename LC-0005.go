func palindromeExpand(s string, l int, r int) string {
    for l >= 0 && r < len(s) && s[l] == s[r] {
        l--
        r++
    }
    return s[l+1:r]
}
func longestPalindrome(s string) string {
    ret := ""
    for i, _ := range s {
        candidate := palindromeExpand(s, i, i)
        if len(candidate) > len(ret) {
            ret = candidate
        }
        
        if i + 1 < len(s) && s[i] == s[i+1] {
            candidate = palindromeExpand(s, i, i + 1)
            if len(candidate) > len(ret) {
                ret = candidate
            }
        }
    }
    
    return ret
}