---
title: "Installing Go"
description: "Install Go, check the version, and know GOROOT vs GOPATH"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

Download Go from [go.dev/dl](https://go.dev/dl) or use your package manager, then check that `go` is on your `PATH`.

## Linux install

```sh
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.xx.x.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin   # add to ~/.profile or ~/.zshrc
go version
```

macOS and Windows have installers on the same page.

## Paths Go uses

| Variable | Default | Holds |
| --- | --- | --- |
| `GOROOT` | install dir (`/usr/local/go`) | The Go toolchain itself |
| `GOPATH` | `~/go` | Module cache (`pkg/mod`) and `go install` binaries (`bin`) |

Add `$(go env GOPATH)/bin` to `PATH` so tools you install with `go install` can run. `go env` prints every setting.

## Toolchain versions

A module's `go.mod` declares the minimum Go version (`go 1.22`). Since Go 1.21, if your installed Go is older, the `go` command can download the required toolchain automatically.

:::caution[Pitfalls]
- Don't set `GOROOT` by hand unless you have several installs; a wrong value breaks the toolchain.
- With modules, your code can live anywhere. It does not need to be inside `GOPATH`.
:::
