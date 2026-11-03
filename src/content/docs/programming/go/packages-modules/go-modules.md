---
title: "Go Modules"
description: "go.mod, go.sum and the commands that manage dependencies"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

A module is a tree of packages with a `go.mod` at its root. It declares the module path, the Go version, and its dependencies.

## go.mod

```text title="go.mod"
module github.com/me/myapp

go 1.22

require github.com/gofiber/fiber/v2 v2.52.0
```

`go.sum` records checksums of every dependency so builds are reproducible and tamper-evident. Commit both files.

## Commands

```sh
go mod init github.com/me/myapp   # create go.mod
go get github.com/gofiber/fiber/v2 # add a dependency
go mod tidy                        # add missing, remove unused
go list -m all                     # show the full dependency list
```

Downloads go to the module cache (`$GOPATH/pkg/mod`) through the proxy `proxy.golang.org`.

:::caution[Pitfalls]
- Run `go mod tidy` before committing; CI often fails on an untidy `go.mod`.
- Private repos need `GOPRIVATE=github.com/yourorg/*` so the public proxy isn't used.
:::
