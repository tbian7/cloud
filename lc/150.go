package lc

func jump(nums []int) int {
	// DP
	// steps := make([]int, len(nums))
	// for i := 0; i < len(nums) - 1; i++ {
	//     for j := i + 1; j <= i + nums[i] && j < len(nums); j++ {
	//         if steps[j] == 0 || steps[j] > steps[i] + 1 {
	//             steps[j] = steps[i] + 1
	//         }
	//     }
	// }
	// return steps[len(nums) - 1]

	// Greedy
	level, left, right := 0, 0, 0
	for right < len(nums)-1 {
		nr := 0
		for i := left; i <= right; i++ {
			if i+nums[i] > nr {
				nr = i + nums[i]
			}
		}
		left, right = right+1, nr
		level++
	}
	return level
}
