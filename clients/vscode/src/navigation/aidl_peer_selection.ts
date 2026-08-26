export interface AidlPeerCandidate {
  path: { raw: string; repoPath?: string };
  localPath?: string;
  confidence: string;
  exists: boolean;
}

export function selectAidlPeerPath(result: { peers: AidlPeerCandidate[] }): string | undefined {
  const targets = result.peers
    .filter((peer) => peer.confidence === "exact" && peer.exists)
    .map((peer) => peer.path.repoPath ?? peer.path.raw)
    .filter((repo) => !repo.startsWith("/") && !/^[A-Za-z]:[\\/]/.test(repo) && !repo.startsWith("\\\\"))
    .filter((repo) => repo.length > 0)
    .filter((repo, index, values) => values.indexOf(repo) === index);
  return targets.length === 1 ? targets[0] : undefined;
}

export function selectAidlPeerLocalPath(result: { peers: AidlPeerCandidate[] }): string | undefined {
  const targets = result.peers
    .filter((peer) => peer.confidence === "exact" && peer.exists && peer.localPath)
    .map((peer) => peer.localPath as string)
    .filter((value, index, values) => values.indexOf(value) === index);
  return targets.length === 1 ? targets[0] : undefined;
}
