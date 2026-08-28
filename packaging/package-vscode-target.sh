#!/usr/bin/env bash
set -euo pipefail

repo_root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
target=${1:?usage: package-vscode-target.sh <os-arch> [engine-root]}
engine_root=${2:-$repo_root/dist/engine}
version=${AOSP_IDE_VERSION:-0.1.0-dev}

case "$target" in
  darwin-amd64) marketplace_target=darwin-x64; engine_name=aosp-index ;;
  darwin-arm64) marketplace_target=darwin-arm64; engine_name=aosp-index ;;
  linux-amd64) marketplace_target=linux-x64; engine_name=aosp-index ;;
  linux-arm64) marketplace_target=linux-arm64; engine_name=aosp-index ;;
  windows-amd64) marketplace_target=win32-x64; engine_name=aosp-index.exe ;;
  windows-arm64) marketplace_target=win32-arm64; engine_name=aosp-index.exe ;;
  *) printf 'unsupported VS Code target: %s\n' "$target" >&2; exit 1 ;;
esac

engine="$engine_root/$target/$engine_name"
test -f "$engine"
temp_root=$(mktemp -d)
trap 'rm -rf "$temp_root"' EXIT
stage="$temp_root/extension"
mkdir -p "$stage/out" "$stage/engine/$target" "$stage/resources" "$stage/l10n"
cp "$repo_root/clients/vscode/package.json" "$stage/"
cp "$repo_root/clients/vscode/package.nls.json" "$stage/"
cp "$repo_root/clients/vscode/package.nls.zh-cn.json" "$stage/"
cp "$repo_root/clients/vscode/.vscodeignore" "$stage/"
cp -R "$repo_root/clients/vscode/resources/." "$stage/resources/"
cp -R "$repo_root/clients/vscode/l10n/." "$stage/l10n/"
cp -R "$repo_root/clients/vscode/out/." "$stage/out/"
cp "$engine" "$stage/engine/$target/$engine_name"
chmod +x "$stage/engine/$target/$engine_name" 2>/dev/null || true
mkdir -p "$repo_root/dist/vscode"
output="$repo_root/dist/vscode/aosp-ide-vscode-$version-$target.vsix"
(
  cd "$stage"
  npx --prefix "$repo_root/clients/vscode" vsce package \
    --no-dependencies --allow-missing-repository --skip-license \
    --target "$marketplace_target" --out "$output"
)
printf '%s\n' "$output"
