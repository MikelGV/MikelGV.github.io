# syntax=docker/dockerfile:1

# Stage 1 — Build
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Download Tailwind CSS v4 standalone CLI (musl build for Alpine)
RUN apk add --no-cache curl libstdc++ libgcc \
 && curl -fsSL -o /usr/local/bin/tailwindcss https://github.com/tailwindlabs/tailwindcss/releases/download/v4.1.4/tailwindcss-linux-x64-musl \
 && chmod +x /usr/local/bin/tailwindcss

# Download module dependencies (layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Generate templ code, build minified CSS, compile binary
RUN go tool templ generate \
 && tailwindcss -i static/css/input.css -o static/css/output.css --minify \
 && CGO_ENABLED=0 go build -o /app/server ./cmd/web

# Stage 2 — Runtime
FROM alpine:latest AS runtime

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server .
COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["./server"]
