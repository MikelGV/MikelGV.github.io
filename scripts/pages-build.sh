#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."

# 1. Generate templ code from *.templ files.
go tool templ generate

# 2. Build Tailwind CSS with the same standalone CLI version as Docker/Vercel.
TAILWIND_VERSION="v4.1.4"
TAILWIND_BIN="${TMPDIR:-/tmp}/tailwindcss-${TAILWIND_VERSION}"
if [[ ! -x "$TAILWIND_BIN" ]]; then
  curl -fsSL --retry 3 -o "$TAILWIND_BIN" \
    "https://github.com/tailwindlabs/tailwindcss/releases/download/${TAILWIND_VERSION}/tailwindcss-linux-x64"
  chmod +x "$TAILWIND_BIN"
fi
"$TAILWIND_BIN" -i static/css/input.css -o static/css/output.css --minify

# 3. Compile everything, then export the static Pages site.
go build ./...
go run ./cmd/export -out dist
