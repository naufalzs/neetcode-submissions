func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	for i := 0 ; i < len(nums); i++ {
		if i == 0 {
			left[i] = 1
			right[len(nums)-1-i] = 1
		} else {
			left[i] = left[i-1] * nums[i-1]
			right[len(nums)-1-i] = right[len(nums)-i] * nums[len(nums)-i]
		}
	}
	res := make([]int, len(nums))
	for i := 0 ; i < len(left); i++ {
		res[i] = left[i] * right[i]
	}
	return res
}
