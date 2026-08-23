package aidl

import (
	"path/filepath"
	"testing"
)

func TestLoadTreeBuildsAidlIndex(t *testing.T) {
	report, err := LoadTree(filepath.Join("..", "..", "testdata", "aidl", "source"))
	if err != nil {
		t.Fatal(err)
	}
	if len(report.Index.Documents) != 3 {
		t.Fatalf("documents: %+v", report.Index.Documents)
	}
	if resolution := report.Index.Resolve(report.Index.Documents["android/demo/IFoo.aidl"], "android.other.IBar"); resolution.State != ResolutionFound {
		t.Fatalf("resolution: %+v diagnostics=%+v", resolution, report.Diagnostics)
	}
}
