---
title: "Cross-Compilation"
description: "Build for another OS and CPU with GOOS and GOARCH"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

Go cross-compiles out of the box: set `GOOS` and `GOARCH` and build. No extra toolchain needed for pure-Go code.

## Commands

```sh
GOOS=linux   GOARCH=amd64 go build -o app-linux .
GOOS=linux   GOARCH=arm64 go build -o app-arm64 .   # Graviton, Raspberry Pi
GOOS=darwin  GOARCH=arm64 go build -o app-mac .
GOOS=windows GOARCH=amd64 go build -o app.exe .
go tool dist list                                   # every supported pair
```

## Production builds

```sh
CGO_ENABLED=0 go build -trimpath -ldflags="-s -w -X main.version=1.4.2" -o app ./cmd/api
```

| Flag | Effect |
| --- | --- |
| `CGO_ENABLED=0` | Static binary, runs on `scratch`/distroless images |
| `-trimpath` | Remove local paths from the binary |
| `-ldflags="-s -w"` | Strip debug info: smaller binary |
| `-X main.version=...` | Set a string variable at build time |

:::caution[Pitfalls]
- Packages using cgo can't cross-compile this simply; you need a C cross-compiler or a pure-Go alternative.
:::
