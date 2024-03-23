package lc

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	var line, res []string
	var ll int
	for _, word := range words {
		if len(word)+len(line)+ll > maxWidth {
			extraSpace := maxWidth - ll
			wc := 1 // words that need to append space to
			if len(line) > 1 {
				wc = len(line) - 1
			}
			space := extraSpace / wc
			remainder := extraSpace % wc
			for i := 0; i < wc; i++ {
				sc := space
				if remainder > 0 {
					sc++
					remainder--
				}
				line[i] = line[i] + strings.Repeat(" ", sc)
			}
			res = append(res, strings.Join(line, ""))
			line, ll = line[:0], 0
		}

		line = append(line, word)
		ll += len(word)
	}
	lastLine := strings.Join(line, " ")
	lastLine += strings.Repeat(" ", maxWidth-len(lastLine))
	return append(res, lastLine)
}
