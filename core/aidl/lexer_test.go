package aidl

import "testing"

func TestLexSkipsCommentsAndKeepsRanges(t *testing.T) {
	tokens, diagnostics := Lex("// comment\npackage demo; /* x */")
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics: %+v", diagnostics)
	}
	if tokens[0].Text != "package" || tokens[0].Start.Line != 2 {
		t.Fatalf("tokens: %+v", tokens)
	}
}
