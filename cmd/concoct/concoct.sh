#!/usr/bin/env bash
set -euo pipefail

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
repository_root="$(cd "$script_dir/../.." && pwd)"

# Compatibility entry point for source checkouts. Distribution builds should
# install the Go binary as `concoct` directly.
CONCOCT_CALLER_DIR="$PWD" exec go -C "$repository_root" run ./cmd/concoct "$@"
