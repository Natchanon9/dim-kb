---
title: "History of Go"
description: "Where Go came from and the releases that changed how it is written"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

Go started at Google in 2007, designed by Robert Griesemer, Rob Pike and Ken Thompson. It was announced as open source in November 2009.

## Why Google built it

Large C++ services took a long time to build, and big codebases were hard to read. The goal was a language with fast builds, simple syntax, and good support for multicore machines and networked servers.

## Releases worth knowing

| Version | Year | What changed |
| --- | --- | --- |
| 1.0 | 2012 | First stable release and the **Go 1 compatibility promise** |
| 1.5 | 2015 | Compiler and runtime rewritten in Go; low-latency concurrent GC |
| 1.11 | 2018 | **Go modules** (dependency management without GOPATH) |
| 1.13 | 2019 | Error wrapping: `%w`, `errors.Is`, `errors.As` |
| 1.18 | 2022 | **Generics** |
| 1.21 | 2023 | `min`/`max` builtins, `slices` and `maps` packages, `log/slog` |
| 1.22 | 2024 | Per-iteration loop variables, `for i := range n`, method + path patterns in `ServeMux` |
| 1.23 | 2024 | Range over functions (iterators) |

## The compatibility promise

Code written for Go 1.0 still compiles with today's Go 1.x. That is why old tutorials mostly still work, and why language changes arrive slowly.

:::caution[Pitfalls]
- Old tutorials may use GOPATH mode, `ioutil`, or capture loop variables manually. These are outdated since modules (1.11), Go 1.16, and Go 1.22.
:::
