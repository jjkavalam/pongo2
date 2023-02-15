package pongo2

import (
	"testing"
)

func Test_dedent(t *testing.T) {
	testDedent(t, "\n b", 1, "\nb")
	testDedent(t, "\n b \n c", 1, "\nb \nc")
	testDedent(t, "\n b \nc", 1, "\nb \nc")
	testDedent(t, "\nb\n", 1, "\nb\n")
}

func testDedent(t *testing.T, s string, l int, expected string) {
	t.Run(s, func(t *testing.T) {
		out := dedentAfterNewline(s, l)
		if out != expected {
			t.Errorf("got '%s'; want '%s'", out, expected)
		}
	})
}
