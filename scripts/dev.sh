#!/bin/sh
set -e

cleanup() {
  kill $(jobs -p) 2>/dev/null
  wait
}

trap cleanup EXIT INT TERM

echo "Starting templ generate watch..."
go tool templ generate -watch &

echo "Starting Tailwind CSS watch..."
tailwindcss -i static/css/input.css -o static/css/output.css --watch &

echo "Starting templ proxy (port 7331 -> 8080)..."
go tool templ proxy --port 7331 --proxy-to 8080 &

sleep 2

echo "Starting air (Go hot reload)..."
air &

wait
