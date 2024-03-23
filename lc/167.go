package lc

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	res := make([]int, 2)

	for left < right {
		if sum := numbers[left] + numbers[right]; sum == target {
			res[0], res[1] = left+1, right+1
			break
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return res
}
