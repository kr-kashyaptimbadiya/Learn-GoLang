func setZeroes(matrix [][]int) {
	r := make([]int, 0)
	c := make([]int, 0)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				r = append(r, i)
				c = append(c, j)
			}
		}
	}

	for _, vr := range r {
		for vc := 0; vc < len(matrix[0]); vc++ {
			matrix[vr][vc] = 0
		}
	}
	for _, vc := range c {
		for vr := 0; vr < len(matrix); vr++ {
			matrix[vr][vc] = 0
		}
	}
}