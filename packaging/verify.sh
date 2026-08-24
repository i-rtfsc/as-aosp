#!/usr/bin/env bash
set -euo pipefail

repo_root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
dist_root="$repo_root/dist"
manifest="$dist_root/manifest.sha256"
test -f "$manifest"
(cd "$dist_root" && shasum -a 256 -c manifest.sha256)
test "$(find "$dist_root/engine" -type f -name 'aosp-index' | wc -l | tr -d ' ')" -eq 4
test "$(find "$dist_root/engine" -type f -name 'aosp-index.exe' | wc -l | tr -d ' ')" -eq 2
test "$(find "$dist_root/jetbrains" -type f -name '*.zip' | wc -l | tr -d ' ')" -ge 1
test "$(find "$dist_root/vscode" -type f -name '*.vsix' | wc -l | tr -d ' ')" -ge 1
printf 'AOSP IDE artifacts verified\n'
