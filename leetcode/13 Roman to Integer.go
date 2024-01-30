func romanToInt(s string) int {
	m := make(map[byte]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000

	// var m = map[byte]int{
	// 'I' : 1, 'V' : 5, 'X' : 10, 'L' : 50, 'C' : 100,'D' : 500, 'M' : 1000,
	// }

	ans := m[s[len(s)-1]]
	for i := len(s) - 2; i >= 0; i-- {
		if m[s[i]] < m[s[i+1]] {
			ans -= m[s[i]]
		} else {
			ans += m[s[i]]
		}
	}
	return ans

}