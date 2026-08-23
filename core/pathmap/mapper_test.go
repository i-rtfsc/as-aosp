package pathmap

import (
	"testing"

	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/topology"
)

func TestMapperConvertsViewAndRepoPaths(t *testing.T) {
	mapper := Mapper{Topology: topology.WorkspaceTopology{SourceMounts: []topology.SourceMount{{
		RepoPrefix: "frameworks/base", ViewPath: "frameworks/base", LocalPath: "/full/frameworks/base", Priority: 100,
	}}}}
	repo, err := mapper.ViewToRepo("frameworks/base/core/java")
	if err != nil || repo != "frameworks/base/core/java" {
		t.Fatalf("repo=%q err=%v", repo, err)
	}
	local, err := mapper.RepoToLocal("frameworks/base/core/java")
	if err != nil || local != "/full/frameworks/base/core/java" {
		t.Fatalf("local=%q err=%v", local, err)
	}
}

func TestMapperConvertsExternalPath(t *testing.T) {
	mapper := Mapper{Mappings: []model.PathMapping{{ExternalPrefix: "/srv/aosp", LocalPrefix: "/work/aosp"}}}
	local, ok := mapper.ExternalToLocal("/srv/aosp/frameworks/base")
	if !ok || local != "/work/aosp/frameworks/base" {
		t.Fatalf("local=%q ok=%v", local, ok)
	}
}

func TestMapperConvertsLocalPath(t *testing.T) {
	mapper := Mapper{Topology: topology.WorkspaceTopology{SourceMounts: []topology.SourceMount{{
		RepoPrefix: "system/core", ViewPath: "core", LocalPath: "/full/system/core", Priority: 100,
	}}}}
	repo, err := mapper.LocalToRepo("/full/system/core/init/init.cpp")
	if err != nil || repo != "system/core/init/init.cpp" {
		t.Fatalf("repo=%q err=%v", repo, err)
	}
}

func TestMapperUsesLongestExternalMapping(t *testing.T) {
	mapper := Mapper{Mappings: []model.PathMapping{
		{ExternalPrefix: "/srv", LocalPrefix: "/short"},
		{ExternalPrefix: "/srv/aosp", LocalPrefix: "/long"},
	}}
	local, ok := mapper.ExternalToLocal("/srv/aosp/system/core")
	if !ok || local != "/long/system/core" {
		t.Fatalf("local=%q ok=%v", local, ok)
	}
}
