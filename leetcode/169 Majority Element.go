func majorityElement(nums []int) int {
	cnt := 0
	var ans int
	for i := 0; i < len(nums); i++ {
		if cnt == 0 {
			ans = nums[i]
		}
		if nums[i] == ans {
			cnt++
		} else {
			cnt--
		}
	}
	return ans
}