#!/usr/bin/env bash
set -euo pipefail

repo_root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
dist_root="$repo_root/dist"
version="${AOSP_IDE_VERSION:-0.1.0-dev}"
rm -rf "$dist_root"
mkdir -p "$dist_root/engine" "$dist_root/jetbrains" "$dist_root/vscode"

for target in darwin-amd64 darwin-arm64 linux-amd64 linux-arm64 windows-amd64 windows-arm64; do
  os=${target%-*}
  arch=${target#*-}
  output="$dist_root/engine/$target/aosp-index"
  mkdir -p "$(dirname "$output")"
  if [[ "$os" == windows ]]; then output="$output.exe"; fi
  case "$arch" in amd64) go_arch=amd64 ;; arm64) go_arch=arm64 ;; esac
  case "$os" in darwin) go_os=darwin ;; linux) go_os=linux ;; windows) go_os=windows ;; esac
  CGO_ENABLED=0 GOOS="$go_os" GOARCH="$go_arch" go build -trimpath -o "$output" "$repo_root/cmd/aosp-index"
done

(
  cd "$repo_root/clients/jetbrains"
  ./gradlew --no-daemon buildPlugin
  plugin_zip=$(find build/distributions -maxdepth 1 -type f -name '*.zip' | head -1)
  test -n "$plugin_zip"
  staging=$(mktemp -d)
  trap 'rm -rf "$staging"' EXIT
  unzip -q "$plugin_zip" -d "$staging"
  plugin_root=$(find "$staging" -mindepth 1 -maxdepth 1 -type d | head -1)
  mkdir -p "$plugin_root/engine"
  cp -R "$dist_root/engine/." "$plugin_root/engine/"
  rebuilt="$dist_root/jetbrains/$(basename "$plugin_zip")"
  (cd "$staging" && zip -qr "$rebuilt" .)
)
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
