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
  cp build/distributions/*.zip "$dist_root/jetbrains/"
)
(
  cd "$repo_root/clients/vscode"
  npm ci
  npm run package
  cp ./*.vsix "$dist_root/vscode/"
)

printf 'AOSP IDE %s packaging complete\n' "$version"
