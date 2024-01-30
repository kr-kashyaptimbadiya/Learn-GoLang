
import "math"

func differenceOfSum(nums []int) int {
	ans := 0
	ans1 := 0
	digit := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		ans += nums[i]
	}
	for i := 0; i < len(nums); i++ {
		for nums[i] != 0 {
			temp := nums[i] % 10
			digit = append(digit, temp)
			nums[i] /= 10
		}
	}
	for i := 0; i < len(digit); i++ {
		ans1 += digit[i]
	}

	final := int(math.Abs(float64(ans - ans1)))
	return final
}