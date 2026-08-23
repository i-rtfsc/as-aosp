package aidl

import (
	"strings"
	"testing"

	"github.com/i-rtfsc/as-aosp/core/metadata"
)

func TestGeneratedPeersInferJavaCppAndNdkPaths(t *testing.T) {
	document := Parse("android/demo/IFoo.aidl", "package android.demo; interface IFoo {}")
	report := NewIndex().GeneratedPeers(document, "android.demo.IFoo", PeerOptions{})
	if len(report.Diagnostics) != 0 || len(report.Peers) != 6 {
		t.Fatalf("report: %+v", report)
	}
	if report.Peers[0].Confidence != "inferred" {
		t.Fatalf("peers: %+v", report.Peers)
	}
}

func TestGeneratedPeersUpgradeVisibleBuildPath(t *testing.T) {
	document := Parse("android/demo/IFoo.aidl", "package android.demo; interface IFoo {}")
	generated := metadata.NewPathRef("/out/soong/generated/android/demo/IFoo.java")
	report := NewIndex().GeneratedPeers(document, "android.demo.IFoo", PeerOptions{GeneratedPaths: []metadata.PathRef{generated}, Exists: func(metadata.PathRef) bool { return true }})
	for _, peer := range report.Peers {
		if peer.Language == PeerJava && peer.Confidence != "exact" || peer.Language == PeerJava && !peer.Exists {
			t.Fatalf("java peer: %+v", peer)
		}
	}
}

func TestGeneratedPeersDoNotCreateNestedJavaFile(t *testing.T) {
	document := Parse("android/demo/Outer.aidl", "package android.demo; parcelable Outer { enum Kind { ONE } }")
	report := NewIndex().GeneratedPeers(document, "android.demo.Outer.Kind", PeerOptions{})
	for _, peer := range report.Peers {
		if peer.Language == PeerJava && peer.Path.Raw != "android/demo/Outer.java" {
			t.Fatalf("peer: %+v", peer)
		}
	}
}

func TestGeneratedPeersWarnWhenMetadataPathIsMissing(t *testing.T) {
	document := Parse(
		"android/demo/IFoo.aidl",
		"package android.demo; interface IFoo {}",
	)
	generated := metadata.NewPathRef(
		"/out/generated/android/demo/IFoo.java",
	)
	report := NewIndex().GeneratedPeers(
		document,
		"android.demo.IFoo",
		PeerOptions{
			GeneratedPaths: []metadata.PathRef{generated},
			Exists:         func(metadata.PathRef) bool { return false },
		},
	)
	if len(report.Diagnostics) != 1 ||
		report.Diagnostics[0].Code != "AIDL_GENERATED_MISSING" {
		t.Fatalf("report: %+v", report)
	}
}

func TestGeneratedPeersUseSrcJarEntry(t *testing.T) {
	document := Parse(
		"android/demo/IFoo.aidl",
		"package android.demo; interface IFoo {}",
	)
	archive := metadata.NewPathRef("/out/generated/framework.srcjar")
	report := NewIndex().GeneratedPeers(
		document,
		"android.demo.IFoo",
		PeerOptions{
			SourceArchives: []metadata.PathRef{archive},
			ArchiveEntries: []ArchiveEntry{{
				Archive:   archive,
				Entry:     "android/demo/IFoo.java",
				LocalPath: "/cache/android/demo/IFoo.java",
			}},
		},
	)
	for _, peer := range report.Peers {
		if peer.Language == PeerJava {
			if peer.Origin != "build-srcjar" ||
				peer.Confidence != "exact" ||
				peer.LocalPath == "" {
				t.Fatalf("peer: %+v", peer)
			}
			return
		}
	}
	t.Fatal("Java peer not found")
}

func TestGeneratedPeersKeepCppAndNdkHeadersSeparate(t *testing.T) {
	document := Parse(
		"android/demo/IFoo.aidl",
		"package android.demo; interface IFoo {}",
	)
	report := NewIndex().GeneratedPeers(
		document,
		"android.demo.IFoo",
		PeerOptions{
			GeneratedPaths: []metadata.PathRef{
				metadata.NewPathRef("/out/android/demo/BpFoo.h"),
				metadata.NewPathRef("/out/aidl/android/demo/BpFoo.h"),
			},
			Exists: func(metadata.PathRef) bool { return true },
		},
	)
	for _, peer := range report.Peers {
		if peer.Kind != PeerHeader ||
			!strings.HasSuffix(peer.Path.Raw, "BpFoo.h") {
			continue
		}
		if peer.Language == PeerCpp &&
			strings.Contains(peer.Path.Raw, "/aidl/") {
			t.Fatalf("C++ peer used NDK path: %+v", peer)
		}
		if peer.Language == PeerNdk &&
			!strings.Contains(peer.Path.Raw, "/aidl/") {
			t.Fatalf("NDK peer used C++ path: %+v", peer)
		}
	}
}
