#!/usr/bin/env bash
set -euo pipefail

repo_root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
dist_root="$repo_root/dist"
manifest="$dist_root/manifest.sha256"
json_manifest="$dist_root/manifest.json"
test -f "$manifest"
test -f "$json_manifest"
(cd "$dist_root" && shasum -a 256 -c manifest.sha256)
node -e '
const value=require(process.argv[1]);
if(!value.version||value.artifacts.length<8) process.exit(1);
const engines=value.artifacts.filter((item)=>item.kind==="engine");
const matrix=new Set(engines.map((item)=>`${item.os}-${item.arch}`));
if(engines.length!==6||matrix.size!==6) process.exit(1);
if(engines.some((item)=>!item.executable)) process.exit(1);
' "$json_manifest"
test "$(find "$dist_root/engine" -type f -name 'aosp-index' | wc -l | tr -d ' ')" -eq 4
test "$(find "$dist_root/engine" -type f -name 'aosp-index.exe' | wc -l | tr -d ' ')" -eq 2
test "$(find "$dist_root/jetbrains" -type f -name '*.zip' | wc -l | tr -d ' ')" -ge 1
test "$(find "$dist_root/vscode" -type f -name '*.vsix' | wc -l | tr -d ' ')" -eq 6
jetbrains_zip=$(find "$dist_root/jetbrains" -type f -name '*.zip' | head -1)
test "$(unzip -Z1 "$jetbrains_zip" | grep -c '/engine/darwin-arm64/aosp-index$')" -eq 1
test "$(unzip -Z1 "$jetbrains_zip" | grep -c '/engine/windows-amd64/aosp-index.exe$')" -eq 1
for target in darwin-amd64 darwin-arm64 linux-amd64 linux-arm64 windows-amd64 windows-arm64; do
  vsix="$dist_root/vscode/aosp-ide-vscode-${AOSP_IDE_VERSION:-0.1.0-dev}-$target.vsix"
  engine_name=aosp-index
  [[ "$target" == windows-* ]] && engine_name=aosp-index.exe
  test "$(unzip -Z1 "$vsix" | grep -c "^extension/engine/$target/$engine_name$")" -eq 1
done
printf 'AOSP Workspace artifacts verified\n'
