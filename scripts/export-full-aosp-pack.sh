#!/usr/bin/env bash
set -Eeuo pipefail

# Export every available build fact from a completed or partially completed AOSP out.
# This script intentionally does not pass --module, so the pack is not module-filtered.

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
AOSP_ROOT=${AOSP_ROOT:-/home/solo/code/aosp}
OUT_ROOT=${OUT_ROOT:-$AOSP_ROOT/out}
PACK_BASE=${PACK_BASE:-/home/solo/code/aosp-packs}
AOSP_PACK_NAME=${AOSP_PACK_NAME:-android-17-cp2a-aosp_cf_x86_64_phone-full}
PACK_ROOT=${PACK_ROOT:-$PACK_BASE/$AOSP_PACK_NAME}
ZIP_ROOT=${ZIP_ROOT:-$PACK_BASE/$AOSP_PACK_NAME.zip}
ENGINE="$SCRIPT_DIR/aosp-index"

log() { printf '[aosp-pack] %s\n' "$*"; }
fail() { printf '[aosp-pack] ERROR: %s\n' "$*" >&2; exit 1; }

[[ -d "$AOSP_ROOT" ]] || fail "AOSP root does not exist: $AOSP_ROOT"
[[ -d "$OUT_ROOT" ]] || fail "out root does not exist: $OUT_ROOT"
[[ -x "$ENGINE" ]] || fail "aosp-index is missing or not executable: $ENGINE"

log "using engine $ENGINE"
log "engine sha256 $(sha256sum "$ENGINE" | awk '{print $1}')"

if [[ "${ALLOW_ACTIVE_BUILD:-0}" != "1" ]] &&
  pgrep -af '(soong_ui\.bash|/ninja |/siso )' | grep -F "$AOSP_ROOT" >/dev/null 2>&1; then
  fail "an AOSP build is still running; wait for it to finish or set ALLOW_ACTIVE_BUILD=1"
fi

cd "$AOSP_ROOT"
source build/envsetup.sh >/dev/null
lunch aosp_cf_x86_64_phone cp2a eng >/dev/null

export ANDROID_BUILD_TOP="$AOSP_ROOT"
export OUT_DIR="$OUT_ROOT"
export SOONG_GEN_COMPDB=1

module_info="$OUT_ROOT/target/product/vsoc_x86_64/module-info.json"
if [[ ! -f "$module_info" ]]; then
  log "module-info.json is missing; attempting incremental m module-info"
  if ! m module-info; then
    fail "m module-info failed; inspect the AOSP build log before exporting"
  fi
fi

log "running full-out preflight"
preflight="$PACK_ROOT.preflight.json"
mkdir -p "$(dirname "$PACK_ROOT")"
"$ENGINE" pack preflight \
  --out-root "$OUT_ROOT" \
  --native > "$preflight"

if command -v jq >/dev/null 2>&1; then
  jq '.' "$preflight" | sed -n '1,180p'
  jq -e '.ready == true' "$preflight" >/dev/null || {
    jq -r '.blockingDiagnostics[]? // empty' "$preflight" >&2
    fail "AOSP out is not ready for a complete Metadata Pack"
  }
else
  cat "$preflight"
fi

log "exporting complete pack to $PACK_ROOT"
"$ENGINE" pack export \
  --aosp-root "$AOSP_ROOT" \
  --out-root "$OUT_ROOT" \
  --destination "$PACK_ROOT" \
  --native

log "validating exported pack"
"$ENGINE" pack validate --root "$PACK_ROOT"

log "inspecting exported pack"
"$ENGINE" pack inspect --root "$PACK_ROOT"

rm -f "$ZIP_ROOT" "$ZIP_ROOT.sha256"
log "creating archive $ZIP_ROOT"
(
  cd "$(dirname "$PACK_ROOT")"
  zip -qr "$ZIP_ROOT" "$(basename "$PACK_ROOT")"
)
sha256sum "$ZIP_ROOT" | tee "$ZIP_ROOT.sha256"
du -sh "$PACK_ROOT" "$ZIP_ROOT"
log "complete"
