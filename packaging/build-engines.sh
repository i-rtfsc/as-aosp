#!/usr/bin/env bash
set -euo pipefail

repo_root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
engine_root=${1:-$repo_root/dist/engine}

mkdir -p "$engine_root"
for target in \
  darwin-amd64 darwin-arm64 \
  linux-amd64 linux-arm64 \
  windows-amd64 windows-arm64; do
  os=${target%-*}
  arch=${target#*-}
  output="$engine_root/$target/aosp-index"
  mkdir -p "$(dirname "$output")"
  if [[ "$os" == windows ]]; then output="$output.exe"; fi
  CGO_ENABLED=0 GOOS="$os" GOARCH="$arch" go build -trimpath \
    -o "$output" "$repo_root/cmd/aosp-index"
done

printf '%s\n' "$engine_root"
