#!/usr/bin/env bash
set -Eeuo pipefail

OUT_ROOT=${1:?usage: migrate-repacked-out-v1.sh <out-root>}
PRODUCT_OUT_NAME=${PRODUCT_OUT_NAME:-vsoc_x86_64}
BACKUP_ROOT=${BACKUP_ROOT:-$(dirname "$OUT_ROOT")/out-v1-private-backup}
DRY_RUN=${DRY_RUN:-0}

log() { printf '[out-migrate] %s\n' "$*"; }
fail() { printf '[out-migrate] ERROR: %s\n' "$*" >&2; exit 1; }
run() { if [[ "$DRY_RUN" == "1" ]]; then printf '+ '; printf '%q ' "$@"; printf '\n'; else "$@"; fi; }

[[ -d "$OUT_ROOT" ]] || fail "out root does not exist: $OUT_ROOT"
[[ -f "$OUT_ROOT/generated/manifest.json" ]] || fail "v1 generated manifest is missing"
[[ -f "$OUT_ROOT/manifest.json" ]] || fail "v1 pack manifest is missing"
command -v jq >/dev/null 2>&1 || fail "jq is required"

move_file() {
  local source=$1 target=$2
  [[ -f "$source" ]] || return 0
  if [[ -e "$target" ]]; then
    cmp -s "$source" "$target" || fail "destination exists with different content: $target"
    run rm -f "$source"
    return 0
  fi
  run mkdir -p "$(dirname "$target")"
  run mv "$source" "$target"
}

log "restoring core metadata paths"
move_file "$OUT_ROOT/module_bp_java_deps.json" "$OUT_ROOT/soong/module_bp_java_deps.json"
move_file "$OUT_ROOT/module_bp_cc_deps.json" "$OUT_ROOT/soong/module_bp_cc_deps.json"
move_file "$OUT_ROOT/module-info.json" "$OUT_ROOT/target/product/$PRODUCT_OUT_NAME/module-info.json"
move_file "$OUT_ROOT/development/ide/compdb/compile_commands.json" \
  "$OUT_ROOT/soong/development/ide/compdb/compile_commands.json"

log "restoring jar paths"
if [[ -d "$OUT_ROOT/out" ]]; then
  while IFS= read -r -d '' source; do
    relative=${source#"$OUT_ROOT/out/"}
    move_file "$source" "$OUT_ROOT/$relative"
  done < <(find "$OUT_ROOT/out" -type f -print0)
fi

log "restoring srcjar paths from generated manifest"
while IFS=$'\t' read -r packed source_relative; do
  [[ -n "$packed" && -n "$source_relative" ]] || continue
  move_file "$OUT_ROOT/$packed" "$OUT_ROOT/$source_relative"
done < <(jq -r '.[] | [.path, .source] | @tsv' "$OUT_ROOT/generated/manifest.json")

if [[ "$DRY_RUN" == "1" ]]; then
  log "dry run complete"
  exit 0
fi

log "verifying restored out"
for required in \
  "$OUT_ROOT/soong/module_bp_java_deps.json" \
  "$OUT_ROOT/soong/module_bp_cc_deps.json" \
  "$OUT_ROOT/soong/development/ide/compdb/compile_commands.json" \
  "$OUT_ROOT/target/product/$PRODUCT_OUT_NAME/module-info.json"; do
  [[ -f "$required" ]] || fail "required restored file is missing: $required"
done

remaining=$(find "$OUT_ROOT/generated/srcjars" -type f 2>/dev/null | wc -l | tr -d ' ')
[[ "$remaining" == "0" ]] || fail "$remaining generated/srcjars files were not migrated"

srcjars=$(find "$OUT_ROOT/soong/.intermediates" -type f -name '*.srcjar' | wc -l | tr -d ' ')
jars=$(find "$OUT_ROOT/soong/.intermediates" -type f -name '*.jar' | wc -l | tr -d ' ')
[[ "$srcjars" -gt 0 ]] || fail "no srcjars found after migration"
[[ "$jars" -gt 0 ]] || fail "no jars found after migration"

log "moving v1 private indexes to $BACKUP_ROOT"
run mkdir -p "$BACKUP_ROOT"
for relative in aidl artifacts jni modules generated manifest.json checksums.sha256; do
  [[ -e "$OUT_ROOT/$relative" ]] || continue
  run mv "$OUT_ROOT/$relative" "$BACKUP_ROOT/$relative"
done
[[ ! -d "$OUT_ROOT/out" ]] || run rm -rf "$OUT_ROOT/out"
[[ ! -d "$OUT_ROOT/development" ]] || run rm -rf "$OUT_ROOT/development"

log "complete: $srcjars srcjars, $jars jars"
log "recoverable v1 indexes: $BACKUP_ROOT"
