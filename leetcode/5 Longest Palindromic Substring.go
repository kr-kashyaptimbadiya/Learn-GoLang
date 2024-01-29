func longestPalindrome(s string) string {
	n := len(s)
	if n == 1 {
		return s
	}
	length := 0
	var ans string = ""

	for i := 0; i < n; i++ {
		p := i
		q := i
		for p >= 0 && q < n && s[p] == s[q] {
			if q-p+1 > length {
				length = q - p + 1
				ans = s[p : p+length]
			}
			p--
			q++
		}
		p = i
		q = i + 1
		for p >= 0 && q < n && s[p] == s[q] {
			if q-p+1 > length {
				length = q - p + 1
				ans = s[p : p+length]
			}
			p--
			q++
		}
	}
	return ans
}