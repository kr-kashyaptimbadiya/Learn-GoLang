func buildArray(nums []int) []int {
	//ans := make([]int,len(nums))
	ans := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		//ans[i] = nums[nums[i]]
		ans = append(ans, nums[nums[i]])
	}
	return ans
}
