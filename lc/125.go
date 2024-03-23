package lc

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	var sb strings.Builder
	for _, r := range s {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			sb.WriteRune(unicode.ToLower(r))
		}
	}
	s = sb.String()
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		if s[left] != s[right] {
			return false
		}
	}
	return true
}
