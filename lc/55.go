package lc

func canJump(nums []int) bool {
	// reachable := make([]bool, len(nums))
	// reachable[0] = true

	// for i := 0; i < len(nums); i++ {
	//     if reachable[i] {
	//         for j := 1; j <= nums[i] && i + j < len(nums); j++ {
	//             reachable[i+j] = true
	//         }
	//     }
	// }
	// return reachable[len(nums) - 1]

	var furthestReachable int
	for i := 0; i < len(nums) && i <= furthestReachable; i++ {
		if i+nums[i] > furthestReachable {
			furthestReachable = i + nums[i]
		}
		if furthestReachable >= len(nums)-1 {
			return true
		}
	}
	return false
}
