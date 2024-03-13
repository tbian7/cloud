package lc

func hIndex(citations []int) int {
	// slices.Sort(citations)

	// for i := 0; i < len(citations); i++{
	//     if citations[i] >= len(citations) - i {
	//         return len(citations) - i
	//     }
	// }
	// return 0
	n := len(citations)
	counts := make([]int, n+1)
	for _, c := range citations {
		if c > n {
			c = n
		}
		counts[c]++
	}

	accumulate := 0
	for i := n; i >= 0; i-- {
		accumulate += counts[i]
		if accumulate >= i {
			return i
		}
	}
	return 0
}
