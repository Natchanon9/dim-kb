---
title: "Go Version History"
description: "What each recent release added that you'll actually use"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

Go ships a release every February and August. Each keeps compatibility with Go 1.0 code. The `go` line in `go.mod` decides which language features a module may use.

## Releases that changed everyday code

| Version | Highlights |
| --- | --- |
| 1.18 (2022) | Generics, fuzzing, workspaces |
| 1.19 | Typed atomics (`atomic.Int64`), doc comment syntax |
| 1.20 | `errors.Join`, multiple `%w`, `context.WithCancelCause` |
| 1.21 | `min`/`max`/`clear`, `slices`, `maps`, `cmp`, `log/slog`, toolchain downloads |
| 1.22 (2024) | Per-iteration loop vars, `range` over int, ServeMux patterns, `math/rand/v2` |
| 1.23 | Range-over-func iterators, `iter`, `unique` |
| 1.24 (2025) | Generic type aliases, `omitzero`, `os.Root`, `testing.B.Loop`, Swiss-table maps |
| 1.25 | Container-aware `GOMAXPROCS`, `sync.WaitGroup.Go`, `testing/synctest` |

## Staying current

- Read the release notes at go.dev/doc/devel/release.
- Run `go fix` / `gopls` modernizers to adopt new idioms.
- Bump the `go` line deliberately; it can change semantics (like loop variables).

:::caution[Pitfalls]
- Using a newer standard-library function than your `go` line allows is reported by `go vet` (stdversion).
:::
