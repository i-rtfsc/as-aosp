#!/usr/bin/env bash
set -euo pipefail

repo_root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
version=${AOSP_IDE_VERSION:-0.1.0-dev}
distribution_root="$repo_root/clients/jetbrains/build/distributions"
thin="$distribution_root/aosp-workspace-jetbrains-$version-thin.zip"
java_home=${JAVA_HOME_21:-${JAVA_HOME:-}}
gradle_args=(--no-daemon)
if [[ -n "${INTELLIJ_PLATFORM_PATH:-}" ]]; then
  gradle_args+=("-PintellijPlatformPath=$INTELLIJ_PLATFORM_PATH")
fi

(
  cd "$repo_root/clients/jetbrains"
  if [[ -n "$java_home" ]]; then
    JAVA_HOME="$java_home" ./gradlew "${gradle_args[@]}" buildPlugin
  else
    ./gradlew "${gradle_args[@]}" buildPlugin
  fi
)
test -f "$thin"
test "$(unzip -Z1 "$thin" | grep -c '/engine/' || true)" -eq 0
rm -f \
  "$distribution_root/aosp-ide-jetbrains-$version.zip" \
  "$distribution_root/aosp-workspace-jetbrains-$version.zip"
printf '%s\n' "$thin"
