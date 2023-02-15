package pongo2

import (
	"strings"
)

func dedentHtmlTokens(tokens []*Token, indentLength int) {
	inMacroScope := false
	for _, token := range tokens {
		// for every HTML token in the macro scope,
		// dedentAfterNewline it by the same measure

		if token.Typ == TokenIdentifier && token.Val == "macro" {
			inMacroScope = true
		} else if token.Typ == TokenIdentifier && token.Val == "endmacro" {
			inMacroScope = false
		} else if inMacroScope && token.Typ == TokenHTML {
			token.Val = dedentAfterNewline(token.Val, indentLength)
		}
	}
}

func trimBlocks(tokens []*Token) {
	prevToken := &Token{
		Typ: TokenHTML,
		Val: "\n",
	}
	for _, token := range tokens {
		// TrimBlocks
		if prevToken.Typ != TokenHTML && token.Typ == TokenHTML && prevToken.Val == "%}" {
			token.Val = trimStartingNewline(token.Val)
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
