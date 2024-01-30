func searchInsert(nums []int, target int) int {
	r := 0
	l := len(nums)
	for r < l {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			l = mid
		} else {
			r = mid + 1
		}
	}
	return r
}