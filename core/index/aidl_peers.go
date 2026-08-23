package index

import (
	"github.com/i-rtfsc/as-aosp/core/aidl"
	"github.com/i-rtfsc/as-aosp/core/graph"
	"github.com/i-rtfsc/as-aosp/core/metadata"
	"github.com/i-rtfsc/as-aosp/core/model"
)

func AidlGeneratedPeers(
	document aidl.Document,
	typeName string,
	moduleGraph graph.ModuleGraph,
	moduleName string,
	exists func(metadata.PathRef) bool,
) aidl.PeerReport {
	return AidlGeneratedPeersWithArchives(
		document,
		typeName,
		moduleGraph,
		moduleName,
		exists,
		nil,
	)
}

func AidlGeneratedPeersWithArchives(
	document aidl.Document,
	typeName string,
	moduleGraph graph.ModuleGraph,
	moduleName string,
	exists func(metadata.PathRef) bool,
	archiveEntries []aidl.ArchiveEntry,
) aidl.PeerReport {
	return aidl.NewIndex().GeneratedPeers(
		document,
		typeName,
		aidl.PeerOptions{
			GeneratedPaths: moduleGraph.GeneratedPaths(moduleName),
			IncludeRoots:   moduleGraph.AidlIncludeDirs(moduleName),
			SourceArchives: moduleGraph.SourceArchives(moduleName),
			ArchiveEntries: archiveEntries,
			Exists:         exists,
		},
	)
}

func AidlGeneratedPeersAutoOwner(
	document aidl.Document,
	typeName string,
	moduleGraph graph.ModuleGraph,
	exists func(metadata.PathRef) bool,
) aidl.PeerReport {
	moduleName, found, ambiguous := moduleGraph.AidlOwner(document.Path)
	if ambiguous {
		return ownerDiagnostic(
			document.Path,
			"AIDL_OWNER_AMBIGUOUS",
			"multiple Soong modules own AIDL source",
		)
	}
	if !found {
		return ownerDiagnostic(
			document.Path,
			"AIDL_OWNER_NOT_FOUND",
			"no Soong module owns AIDL source",
		)
	}
	return AidlGeneratedPeers(
		document,
		typeName,
		moduleGraph,
		moduleName,
		exists,
	)
}

func ownerDiagnostic(
	repoPath model.RepoPath,
	code string,
	message string,
) aidl.PeerReport {
	return aidl.PeerReport{
		Peers: make([]aidl.GeneratedPeer, 0),
		Diagnostics: []model.Diagnostic{{
			Code:     code,
			Severity: model.SeverityWarning,
			Stage:    "aidl-generated",
			Message:  message,
			RepoPath: repoPath,
		}},
	}
}
