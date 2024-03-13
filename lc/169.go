package lc

func majorityElement(nums []int) int {
	// m := make(map[int]int)
	// for _, n := range nums {
	//     m[n]++
	//     if m[n] > len(nums) / 2 {
	//         return n
	//     }
	// }
	// return -1

	if len(nums) < 3 {
		return nums[0]
	}
	var l, r int
	for j := 1; j < len(nums); j++ {
		if r < l {
			l, r = j, j
			continue
		}
		if nums[r] == nums[j] {
			r++
			nums[r] = nums[j]
		} else {
			r--
		}
	}
	return nums[l]
}
