export interface JniDefinitionLink {
  java: { path: string; range: { start: number } };
  native: { path: string; range: { start: number } };
  confidence?: string;
}

export function selectJniDefinitionTarget(result: { peers: JniDefinitionLink[] }, currentPath: string): { path: string; offset: number } | undefined {
  const targets = result.peers.flatMap((peer) => {
    if (peer.confidence && peer.confidence !== "exact") return [];
    if (peer.java.path === currentPath) return [{ path: peer.native.path, offset: peer.native.range.start }];
    if (peer.native.path === currentPath) return [{ path: peer.java.path, offset: peer.java.range.start }];
    return [];
  });
  return targets.length === 1 ? targets[0] : undefined;
}
