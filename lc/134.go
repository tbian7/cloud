package lc

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	remain := make([]int, n)
	for i := 0; i < n; i++ {
		remain[i] = gas[i] - cost[i]
	}

	for i := 0; i < n; {
		sum := 0
		j := i
		for ; j < n || j%n != i; j++ {
			sum += remain[j%n]
			if sum < 0 {
				break
			}
		}
		if j >= n && j%n == i {
			return i
		}
		if j%n < i {
			break
		}
		i = j + 1
	}
	return -1
}
