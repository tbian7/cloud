package lc

func rotate(nums []int, k int) {
	if k == 0 {
		return
	}
	k = k % len(nums)
	flip(nums)
	flip(nums[0:k])
	flip(nums[k:])
}

func flip(nums []int) {
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}
