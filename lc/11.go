package lc

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		if height[left] < height[right] {
			if a := height[left] * (right - left); a > maxArea {
				maxArea = a
			}
			left++
		} else if height[left] > height[right] {
			if a := height[right] * (right - left); a > maxArea {
				maxArea = a
			}
			right--
		} else {
			if a := height[right] * (right - left); a > maxArea {
				maxArea = a
			}
			left++
			right--
		}
	}
	return maxArea
}
