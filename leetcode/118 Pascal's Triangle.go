func generate(numRows int) [][]int {
	ans := make([][]int, 0)
	if numRows == 1 {
		ans = append(ans, []int{1})
		return ans
	}
	ans = append(ans, []int{1})
	ans = append(ans, []int{1, 1})
	for i := 2; i < numRows; i++ {
		temp := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				temp[j] = 1
				continue
			}
			temp[j] = ans[i-1][j-1] + ans[i-1][j]
		}
		ans = append(ans, temp)
	}
	return ans
}