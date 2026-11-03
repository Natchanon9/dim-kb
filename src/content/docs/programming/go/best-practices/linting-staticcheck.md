---
title: "Linting & staticcheck"
description: "go vet, staticcheck and golangci-lint in CI"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

Linters catch bugs the compiler allows. Run them locally and in CI.

## Tools

| Tool | Catches |
| --- | --- |
| `go vet` | Printf mistakes, copied locks, unreachable code (ships with Go) |
| `staticcheck` | Bugs, deprecated APIs, simplifications |
| `golangci-lint` | Runs many linters (incl. staticcheck, errcheck, gosec) in one fast pass |

```sh
go vet ./...
go install honnef.co/go/tools/cmd/staticcheck@latest && staticcheck ./...
golangci-lint run
```

## Minimal config

```yaml title=".golangci.yml"
version: "2"
linters:
  enable:
    - errcheck
    - staticcheck
    - gosec
    - revive
```

:::caution[Pitfalls]
- Enabling every linter creates noise people learn to ignore; start small and add rules that catch real bugs.
:::
