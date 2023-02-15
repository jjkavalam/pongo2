package pongo2

import (
	"strings"
)

func dedentHtmlTokens(tokens []*Token, dedentLength int) {
	prevToken := &Token{
		Typ: TokenHTML,
		Val: "\n",
	}
	hasSeenFirstHtmlToken := false

	inMacroScope := false

	for _, token := range tokens {
		if token.Typ == TokenIdentifier && token.Val == "macro" {
			inMacroScope = true
			hasSeenFirstHtmlToken = false
		} else if token.Typ == TokenIdentifier && token.Val == "endmacro" {
			inMacroScope = false
		}

		if inMacroScope && token.Typ == TokenHTML {
			// dedent HTML tokens
			token.Val = dedentAfterNewline(token.Val, dedentLength)

			// in addition, if this is the first HTML token in the macro block; also
			// trim the starting newline
			if prevToken.Typ == TokenSymbol && prevToken.Val == "%}" && !hasSeenFirstHtmlToken {
				token.Val = trimStartingNewline(token.Val)
			}

			hasSeenFirstHtmlToken = true
		}

		prevToken = token

	}
}

func trimStartingNewline(s string) string {
	if strings.HasPrefix(s, "\n") {
		return s[1:]
	}
	return s
}
