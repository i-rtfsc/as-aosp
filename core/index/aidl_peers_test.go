package index

import (
	"testing"

	"github.com/i-rtfsc/as-aosp/core/aidl"
	"github.com/i-rtfsc/as-aosp/core/graph"
	"github.com/i-rtfsc/as-aosp/core/metadata"
)

func TestAidlGeneratedPeersUsesModuleGraphEvidence(t *testing.T) {
	document := aidl.Parse(
		"android/demo/IFoo.aidl",
		"package android.demo; interface IFoo {}",
	)
	moduleGraph := graph.Build(metadata.Snapshot{
		Java: map[string]metadata.JavaModule{
			"framework": {
				Name: "framework",
				AidlIncludeDirs: []metadata.PathRef{
					metadata.NewPathRef("frameworks/base/core/java"),
				},
				SrcJars: []metadata.PathRef{
					metadata.NewPathRef("/srv/aosp/out/generated/framework.srcjar"),
				},
				Sources: []metadata.PathRef{
					metadata.NewPathRef(
						"/srv/aosp/out/generated/android/demo/IFoo.java",
					),
				},
			},
		},
	})
	report := AidlGeneratedPeers(
		document,
		"android.demo.IFoo",
		moduleGraph,
		"framework",
		func(metadata.PathRef) bool { return true },
	)
	if len(report.IncludeRoots) != 1 || len(report.SourceArchives) != 1 {
		t.Fatalf("evidence: %+v", report)
	}
	for _, peer := range report.Peers {
		if peer.Language == aidl.PeerJava {
			if !peer.Exists || peer.Confidence != "exact" {
				t.Fatalf("peer: %+v", peer)
			}
			return
		}
	}
	t.Fatal("Java peer not found")
}
