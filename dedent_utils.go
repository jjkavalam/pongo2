package pongo2

import (
	"strings"
)

func dedentAfterNewline(s string, l int) string {
	result := ""

	for start := 0; start < len(s); {
		// find next newline
		ret := strings.Index(s[start:], "\n")
		if ret == -1 {
			result += s[start:]
			break
		}
		nl := start + ret
		result += s[start : nl+1]

		// skip utmost l spaces after that
		i := nl + 1
		for ; i < len(s) && i <= nl+l; i++ {
			if s[i] == ' ' {
				continue
			} else {
				break
			}
		}

		start = i
	}

	return result
}
