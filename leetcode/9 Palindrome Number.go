func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	f := 0
	l := len(s) - 1

	for f < l {
		if s[f] != s[l] {
			return false
		}
		f++
		l--
	}
	return true
}