func removeDuplicates(nums []int) int {
	ind := 1
	temp := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == temp {
			continue
		} else {
			nums[ind] = nums[i]
			temp = nums[i]
			ind++
		}
	}
	return ind
}