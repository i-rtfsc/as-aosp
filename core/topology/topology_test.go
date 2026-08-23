package topology

import (
	"testing"

	"github.com/i-rtfsc/as-aosp/core/model"
)

func TestResolveUsesLongestPrefixThenPriority(t *testing.T) {
	base := SourceMount{Name: "base", RepoPrefix: "frameworks/base", ViewPath: "frameworks/base", LocalPath: "/a/base", Priority: 100}
	ui := SourceMount{Name: "ui", RepoPrefix: "frameworks/base/packages/SystemUI", ViewPath: "ui", LocalPath: "/vendor/ui", Priority: 1}
	topology := WorkspaceTopology{SourceMounts: []SourceMount{base, ui}}
	resolution := topology.Resolve(model.RepoPath("frameworks/base/packages/SystemUI/src/SystemUI.java"))
	if resolution.State != ProviderActive || resolution.Mount == nil || resolution.Mount.Name != "ui" {
		t.Fatalf("unexpected resolution: %+v", resolution)
	}
}

func TestResolveReportsSamePriorityConflict(t *testing.T) {
	topology := WorkspaceTopology{SourceMounts: []SourceMount{
		{Name: "a", RepoPrefix: "system/core", ViewPath: "core-a", LocalPath: "/a", Priority: 10},
		{Name: "b", RepoPrefix: "system/core", ViewPath: "core-b", LocalPath: "/b", Priority: 10},
	}}
	resolution := topology.Resolve("system/core/init/init.cpp")
	if resolution.State != ProviderConflict {
		t.Fatalf("got %s", resolution.State)
	}
}
