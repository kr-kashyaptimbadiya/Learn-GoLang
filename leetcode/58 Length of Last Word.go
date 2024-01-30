func lengthOfLastWord(s string) int {
	arrayOfStrings := strings.Fields(s)
	k := len(arrayOfStrings)
	ans := len(arrayOfStrings[k-1])
	return ans
}