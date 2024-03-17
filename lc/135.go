package lc

func candy(ratings []int) int {
	n := len(ratings)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		c[i] = 1
	}

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			c[i] = c[i-1] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && c[i+1]+1 > c[i] {
			c[i] = c[i+1] + 1
		}
	}

	var sum int
	for i := 0; i < n; i++ {
		sum += c[i]
	}
	return sum
}
