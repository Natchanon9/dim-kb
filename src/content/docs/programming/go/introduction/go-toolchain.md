---
title: "Go Toolchain"
description: "The go commands you use every day"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

One binary, `go`, builds, tests, formats and manages dependencies. No separate build tool is needed.

## Daily commands

| Command | Does |
| --- | --- |
| `go run .` | Compile and run the package in this directory |
| `go build -o app .` | Compile to a binary |
| `go test ./...` | Run all tests (`-race`, `-cover`, `-run Name`) |
| `go fmt ./...` | Format code (same as `gofmt -w`) |
| `go vet ./...` | Catch suspicious code the compiler allows |
| `go mod init <path>` | Start a module |
| `go mod tidy` | Add missing / remove unused dependencies |
| `go get pkg@v1.2.3` | Add or change a dependency version |
| `go install pkg@latest` | Build a tool into `$GOPATH/bin` |
| `go env` | Show settings such as `GOPATH`, `GOOS` |
| `go doc fmt.Println` | Show documentation |

## `./...`

`./...` means "this directory and every package below it". Most commands accept it.

:::caution[Pitfalls]
- Run `gofmt` and `go vet` in CI. Unformatted code is a review smell in every Go team.
- `go get` no longer installs binaries; use `go install pkg@version`.
:::
