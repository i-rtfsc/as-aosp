#!/usr/bin/env bash

set -euo pipefail

range=${COMMIT_RANGE:-HEAD}

commits=()
if [[ "$range" == *..* ]]; then
  while IFS= read -r commit; do
    commits+=("$commit")
  done < <(git rev-list --reverse "$range")
else
	commits=("$range")
fi

for commit in "${commits[@]}"; do
  git show -s --format=%B "$commit" | awk -v commit="$commit" '
    length($0) > 75 {
      printf "%s:%d exceeds 75 characters: %s\n", commit, NR, $0
      failed = 1
    }
    END { exit failed }
  '

  count=$(git show -s --format=%B "$commit" |
    grep -Ec '^Change-Id: I[0-9a-f]{40}$' || true)
  if [[ "$count" -ne 1 ]]; then
    printf '%s must contain exactly one Change-Id\n' "$commit"
    exit 1
  fi
done
