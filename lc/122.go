package lc

func maxProfit(prices []int) int {
	var mp int
	// for i := 0; i < len(prices); {
	//     var bought int
	//     for ; i < len(prices); i++ {
	//         if i == len(prices) - 1 {
	//           return mp
	//         }
	//         if prices[i] < prices[i+1] {
	//             bought = prices[i]
	//             break
	//         }
	//     }

	//     for i++; i < len(prices); i++ {
	//         if i == len(prices) - 1 || prices[i] > prices[i+1] {
	//             mp += prices[i] - bought
	//             break
	//         }
	//     }
	// }

	var selling bool
	var bought int
	for i := 0; i < len(prices); i++ {
		if selling {
			if i == len(prices)-1 || prices[i] > prices[i+1] {
				mp += prices[i] - bought
				selling = false
			}
		} else {
			if i == len(prices)-1 {
				break
			}
			if prices[i] < prices[i+1] {
				bought = prices[i]
				selling = true
			}
		}

	}
	return mp
}
