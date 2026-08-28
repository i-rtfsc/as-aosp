#!/usr/bin/env bash
set -euo pipefail

if [[ -z "${ANDROID_BUILD_TOP:-}" ]]; then
  echo "ANDROID_BUILD_TOP is required (AOSP source root)" >&2
  exit 2
fi
if [[ -z "${OUT_DIR:-}" ]]; then
  echo "OUT_DIR is required (AOSP out root)" >&2
  exit 2
fi
if [[ -z "${AOSP_PACK_DESTINATION:-}" ]]; then
  echo "AOSP_PACK_DESTINATION is required (Portable Out destination)" >&2
  exit 2
fi

engine="${AOSP_INDEX_BIN:-aosp-index}"
args=(pack export
  --aosp-root "$ANDROID_BUILD_TOP"
  --out-root "$OUT_DIR"
  --destination "$AOSP_PACK_DESTINATION")
if [[ "${AOSP_PACK_NATIVE:-0}" == "1" ]]; then
  args+=(--native)
fi
if [[ "${AOSP_PACK_IDE:-0}" == "1" ]]; then
  args+=(--ide)
fi
if [[ -n "${AOSP_PACK_MODULES:-}" ]]; then
  IFS=',' read -r -a modules <<< "$AOSP_PACK_MODULES"
  for module in "${modules[@]}"; do
    [[ -n "${module//[[:space:]]/}" ]] && args+=(--module "${module// /}")
  done
fi

preflight_args=(pack preflight --out-root "$OUT_DIR")
if [[ "${AOSP_PACK_NATIVE:-0}" == "1" ]]; then
  preflight_args+=(--native)
fi
if [[ "${AOSP_PACK_REQUIRE_IDE:-0}" == "1" ]]; then
  preflight_args+=(--ide)
fi
if [[ -n "${AOSP_PACK_MODULES:-}" ]]; then
  for module in "${modules[@]}"; do
    [[ -n "${module//[[:space:]]/}" ]] && preflight_args+=(--module "${module// /}")
  done
fi
"$engine" "${preflight_args[@]}"

"$engine" "${args[@]}"
validate_args=(pack validate --root "$AOSP_PACK_DESTINATION")
if [[ "${AOSP_PACK_NATIVE:-0}" == "1" ]]; then
  validate_args+=(--native)
fi
if [[ "${AOSP_PACK_REQUIRE_IDE:-0}" == "1" ]]; then
  validate_args+=(--ide)
fi
"$engine" "${validate_args[@]}"
"$engine" pack inspect --root "$AOSP_PACK_DESTINATION"
