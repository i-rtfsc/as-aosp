package model

import "testing"

func TestNormalizeRepoPath(t *testing.T) {
	got, err := NormalizeRepoPath(`frameworks\\base\\core/java`)
	if err != nil {
		t.Fatal(err)
	}
	if got != "frameworks/base/core/java" {
		t.Fatalf("got %q", got)
	}
}

func TestNormalizeRepoPathRejectsAbsoluteAndEscape(t *testing.T) {
	for _, input := range []string{"/frameworks/base", `C:\\aosp\\frameworks`, "../frameworks"} {
		if _, err := NormalizeRepoPath(input); err == nil {
			t.Fatalf("expected %q to be rejected", input)
		}
	}
}
