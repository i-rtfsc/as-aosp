#!/usr/bin/env bash
set -euo pipefail

repo_root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
version=${AOSP_IDE_VERSION:-0.1.0-dev}
engine_root=${1:-$repo_root/dist/engine}
output_root=${2:-$repo_root/dist/jetbrains}
if [[ "$engine_root" != /* ]]; then engine_root="$repo_root/$engine_root"; fi
if [[ "$output_root" != /* ]]; then output_root="$repo_root/$output_root"; fi
base="$repo_root/clients/jetbrains/build/distributions/aosp-workspace-jetbrains-$version-thin.zip"
output="$output_root/aosp-workspace-jetbrains-$version.zip"
java_home=${JAVA_HOME_21:-${JAVA_HOME:-}}
gradle_args=(--no-daemon)
if [[ -n "${INTELLIJ_PLATFORM_PATH:-}" ]]; then
  gradle_args+=("-PintellijPlatformPath=$INTELLIJ_PLATFORM_PATH")
fi

if [[ ! -x "$engine_root/darwin-arm64/aosp-index" ]] ||
  [[ ! -x "$engine_root/windows-amd64/aosp-index.exe" ]]; then
  bash "$repo_root/packaging/build-engines.sh" "$engine_root" >/dev/null
fi

(
  cd "$repo_root/clients/jetbrains"
  if [[ -n "$java_home" ]]; then
    JAVA_HOME="$java_home" ./gradlew "${gradle_args[@]}" buildPlugin
  else
    ./gradlew "${gradle_args[@]}" buildPlugin
  fi
)
test -f "$base"

staging=$(mktemp -d)
trap 'rm -rf "$staging"' EXIT
unzip -q "$base" -d "$staging"
plugin_root=$(find "$staging" -mindepth 1 -maxdepth 1 -type d | head -1)
test -n "$plugin_root"
mkdir -p "$plugin_root/engine" "$output_root"
cp -R "$engine_root/." "$plugin_root/engine/"
rm -f \
  "$output" \
  "$output_root/aosp-ide-jetbrains-$version.zip"
(cd "$staging" && zip -qr "$output" .)

test "$(unzip -Z1 "$output" | grep -c '/engine/darwin-arm64/aosp-index$')" -eq 1
test "$(unzip -Z1 "$output" | grep -c '/engine/windows-amd64/aosp-index.exe$')" -eq 1
printf '%s\n' "$output"
