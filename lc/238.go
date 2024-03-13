package lc

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	result[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		result[i] = result[i+1] * nums[i]
	}
	for i := 1; i < len(nums); i++ {
		nums[i] *= nums[i-1]
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result[i] = result[i+1]
		} else if i == len(nums)-1 {
			result[i] = nums[i-1]
		} else {
			result[i] = nums[i-1] * result[i+1]
		}
	}
	return result
}
