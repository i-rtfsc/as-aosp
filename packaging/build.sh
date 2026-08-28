#!/usr/bin/env bash
set -euo pipefail

repo_root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
dist_root="$repo_root/dist"
version="${AOSP_IDE_VERSION:-0.1.0-dev}"
rm -rf "$dist_root"
mkdir -p "$dist_root/engine" "$dist_root/jetbrains" "$dist_root/vscode"
bash "$repo_root/packaging/build-engines.sh" "$dist_root/engine" >/dev/null
AOSP_IDE_VERSION="$version" bash \
  "$repo_root/packaging/package-jetbrains-full.sh" \
  "$dist_root/engine" "$dist_root/jetbrains" >/dev/null
(
  cd "$repo_root/clients/vscode"
  npm ci
  npm run compile
)
for target in darwin-amd64 darwin-arm64 linux-amd64 linux-arm64 windows-amd64 windows-arm64; do
  AOSP_IDE_VERSION="$version" bash "$repo_root/packaging/package-vscode-target.sh" \
    "$target" "$dist_root/engine"
done

(
  cd "$dist_root"
  find engine jetbrains vscode -type f -print | LC_ALL=C sort | while read -r file; do
    shasum -a 256 "$file"
  done > manifest.sha256
  AOSP_IDE_VERSION="$version" node "$repo_root/packaging/manifest.js" "$dist_root"
)

printf 'AOSP Workspace %s packaging complete\n' "$version"
