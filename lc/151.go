package lc

import "slices"

func reverseWords(s string) string {
	var b []byte
	var i, j int
	for i < len(s) {
		i, j = findNextWord(s, i)
		if i < len(s) {
			for k := j - 1; k >= i; k-- {
				b = append(b, s[k])
			}
			b = append(b, ' ')
			i = j
		}
	}
	if len(b) > 0 {
		b = b[:len(b)-1]
	}
	slices.Reverse(b)
	return string(b)
}

func findNextWord(s string, i int) (int, int) {
	for i < len(s) && s[i] == ' ' {
		i++
	}
	j := i + 1
	for j < len(s) && s[j] != ' ' {
		j++
	}
	return i, j
}
