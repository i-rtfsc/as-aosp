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
temp_root=$(mktemp -d)
trap 'rm -rf "$temp_root"' EXIT
stage="$temp_root/extension"
mkdir -p "$stage/out" "$stage/engine/$target" "$stage/resources"
(
  cd "$repo_root/clients/vscode"
  npm ci
  npm run compile
  npx vsce package --no-dependencies --help >/dev/null
)
cp "$repo_root/clients/vscode/package.json" "$stage/"
cp "$repo_root/clients/vscode/.vscodeignore" "$stage/"
cp -R "$repo_root/clients/vscode/resources/." "$stage/resources/"
cp -R "$repo_root/clients/vscode/out/." "$stage/out/"
CGO_ENABLED=0 GOOS="$platform" GOARCH="$arch" go build -trimpath -o "$stage/engine/$target/$engine_name" "$repo_root/cmd/aosp-index"
chmod +x "$stage/engine/$target/$engine_name" 2>/dev/null || true
mkdir -p "$repo_root/dist/vscode"
(cd "$stage" && npx --prefix "$repo_root/clients/vscode" vsce package --no-dependencies \
  --allow-missing-repository --skip-license \
  --out "$repo_root/dist/vscode/aosp-ide-vscode-$version-$target.vsix")
printf '%s\n' "$repo_root/dist/vscode/aosp-ide-vscode-$version-$target.vsix"
