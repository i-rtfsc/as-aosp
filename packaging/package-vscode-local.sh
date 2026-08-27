#!/usr/bin/env bash
set -euo pipefail

repo_root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
version="${AOSP_IDE_VERSION:-0.1.0-dev}"
case "$(uname -s)" in
  Darwin) platform=darwin ;;
  Linux) platform=linux ;;
  MINGW*|MSYS*|CYGWIN*) platform=windows ;;
  *) echo "unsupported host OS" >&2; exit 1 ;;
esac
case "$(uname -m)" in
  x86_64|amd64) arch=amd64 ;;
  arm64|aarch64) arch=arm64 ;;
  *) echo "unsupported host architecture" >&2; exit 1 ;;
esac
target="$platform-$arch"
engine_name=aosp-index
[[ "$platform" == windows ]] && engine_name=aosp-index.exe
(
  cd "$repo_root/clients/vscode"
  npm ci
  npm run compile
)
engine_root="$repo_root/dist/engine"
mkdir -p "$engine_root/$target"
CGO_ENABLED=0 GOOS="$platform" GOARCH="$arch" go build -trimpath \
  -o "$engine_root/$target/$engine_name" "$repo_root/cmd/aosp-index"
AOSP_IDE_VERSION="$version" bash "$repo_root/packaging/package-vscode-target.sh" \
  "$target" "$engine_root"
