package lc

func lengthOfLastWord(s string) int {
	// s = strings.TrimSpace(s)
	// return len(s) - 1 - strings.LastIndex(s, " ")

	i, length := len(s)-1, 0

	for i >= 0 && s[i] == ' ' {
		i--
	}

	for i >= 0 && s[i] != ' ' {
		length++
		i--
	}
	return length
}
