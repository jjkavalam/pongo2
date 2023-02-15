package pongo2

import (
	"testing"
)

func Test_dedentHtmlTokens(t *testing.T) {
	x := `
{% macro f() %}
    <a>{{ name }}</a>
  <div>world</div>
{% endmacro %}`

	tokens, err := lex("a", x)
	if err != nil {
		t.Fatal(err)
	}

	dedentHtmlTokens(tokens, 4)

	// when you get a run of HTML tokens
	// find runs of HTML tokens and dedentAfterNewline it
	for _, token := range tokens {
		if token.Typ == TokenHTML {
			t.Log(">" + token.Val + "|")
		}
	}

	// TODO: add assertions
}
