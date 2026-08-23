package index

import (
	"github.com/i-rtfsc/as-aosp/core/aidl"
	"github.com/i-rtfsc/as-aosp/core/graph"
	"github.com/i-rtfsc/as-aosp/core/metadata"
)

func AidlGeneratedPeers(
	document aidl.Document,
	typeName string,
	moduleGraph graph.ModuleGraph,
	moduleName string,
	exists func(metadata.PathRef) bool,
) aidl.PeerReport {
	return aidl.NewIndex().GeneratedPeers(
		document,
		typeName,
		aidl.PeerOptions{
			GeneratedPaths: moduleGraph.GeneratedPaths(moduleName),
			IncludeRoots:   moduleGraph.AidlIncludeDirs(moduleName),
			SourceArchives: moduleGraph.SourceArchives(moduleName),
			Exists:         exists,
		},
	)
}
