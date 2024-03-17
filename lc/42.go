package lc

func trap(height []int) int {
	n := len(height)
	l := make([]int, n)
	r := make([]int, n)

	var h int
	for i := 1; i < n; i++ {
		if h < height[i-1] {
			h = height[i-1]
		}
		l[i] = h
	}

	h = 0
	for i := n - 2; i >= 0; i-- {
		if h < height[i+1] {
			h = height[i+1]
		}
		r[i] = h
	}

	var w int
	for i := 1; i < n-1; i++ {
		if l[i] < r[i] {
			if height[i] < l[i] {
				w += l[i] - height[i]
			}
		} else {
			if height[i] < r[i] {
				w += r[i] - height[i]
			}
		}
	}
	return w
}
