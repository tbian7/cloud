package lc

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	var sb strings.Builder
	for i := 0; i < len(strs[0]); i++ {
		for _, s := range strs {
			if i == len(s) || s[i] != strs[0][i] {
				return sb.String()
			}
		}
		sb.WriteByte(strs[0][i])
	}
	return sb.String()
}
