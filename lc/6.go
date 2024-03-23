package lc

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var sb strings.Builder
	increment := 2*numRows - 2
	for r := 0; r < numRows; r++ {
		littleIncrement := increment - 2*r
		for i := r; i < len(s); i += increment {
			sb.WriteByte(s[i])
			if r != 0 && r != numRows-1 && i+littleIncrement < len(s) {
				sb.WriteByte(s[i+littleIncrement])
			}
		}
	}

	return sb.String()
}
