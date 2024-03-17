package lc

func romanToInt(s string) int {
	sv := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var result int
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && sv[rune(s[i])] < sv[rune(s[i+1])] {
			result -= sv[rune(s[i])]
		} else {
			result += sv[rune(s[i])]
		}
	}
	return result
}
