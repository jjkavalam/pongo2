package pongo2

func dedentHtmlTokens(tokens []*Token, dedentLength int) {
	inMacroScope := false
	for _, token := range tokens {
		if token.Typ == TokenIdentifier && token.Val == "macro" {
			inMacroScope = true
		} else if token.Typ == TokenIdentifier && token.Val == "endmacro" {
			inMacroScope = false
		} else if inMacroScope && token.Typ == TokenHTML {
			token.Val = dedentAfterNewline(token.Val, dedentLength)
		}
	}
}
