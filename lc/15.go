package lc

import "slices"

func threeSum(nums []int) [][]int {
	var res [][]int
	slices.Sort(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			if s := nums[left] + nums[right]; s == -nums[i] {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				right--
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if s < -nums[i] {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
