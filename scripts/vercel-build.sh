#!/usr/bin/env bash
set -euo pipefail

# 1. Generate templ code from *.templ files.
go tool templ generate

# 2. Build Tailwind CSS with the same standalone CLI version as Dockerfile.
TAILWIND_VERSION="v4.1.4"
TAILWIND_BIN="${TMPDIR:-/tmp}/tailwindcss-${TAILWIND_VERSION}"
if [[ ! -x "$TAILWIND_BIN" ]]; then
  curl -fsSL --retry 3 -o "$TAILWIND_BIN" \
    "https://github.com/tailwindlabs/tailwindcss/releases/download/${TAILWIND_VERSION}/tailwindcss-linux-x64"
  chmod +x "$TAILWIND_BIN"
fi
"$TAILWIND_BIN" -i static/css/input.css -o static/css/output.css --minify

# 3. Copy static assets into the directory Vercel serves from its CDN.
rm -rf public/static
mkdir -p public/static
cp -R static/. public/static/

# 4. Compile the Go server. Vercel provides VERCEL_OUTPUT_FILE; locally use ./server.
CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" \
  -o "${VERCEL_OUTPUT_FILE:-server}" ./cmd/api
