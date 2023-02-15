package pongo2

func dedentHtmlTokens(tokens []*Token) {
	inMacroScope := false
	firstHtmlTokenSeen := false
	indentLength := 0
	for _, token := range tokens {
		// measure the indent by looking at the first HTML token
		// in the macro.
		// for every HTML token in the macro scope,
		// dedentAfterNewline it by the same measure

		if token.Typ == TokenIdentifier && token.Val == "macro" {
			inMacroScope = true
			firstHtmlTokenSeen = false
		} else if token.Typ == TokenIdentifier && token.Val == "endmacro" {
			inMacroScope = false
		} else if inMacroScope && token.Typ == TokenHTML {
			if !firstHtmlTokenSeen {
				indentLength = measureIndent(token.Val)
				firstHtmlTokenSeen = true
			}
			token.Val = dedentAfterNewline(token.Val, indentLength)
		}
	}
}
