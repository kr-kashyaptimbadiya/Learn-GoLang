func truncateSentence(s string, k int) string {
	arrayOfString := strings.Fields(s)
	var ans = ""
	for i := 0; i < k-1; i++ {
		ans += arrayOfString[i] + " "
	}
	return ans + arrayOfString[k-1]
}