func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	ans := make([]bool, 0)

	for i := 0; i < len(candies); i++ {
		if max < candies[i] {
			max = candies[i]
		}
	}

	for i := 0; i < len(candies); i++ {

		if candies[i]+extraCandies >= max {
			ans = append(ans, true)
		} else {
			ans = append(ans, false)
		}
	}
	return ans
}