---
title: "Go Playground"
description: "Run and share Go code in the browser"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

[go.dev/play](https://go.dev/play) compiles and runs Go in a sandbox. Use it to test a small idea or share a reproducible snippet.

## What it can do

- Pick the Go version (including the dev branch).
- **Share** gives a permanent link, handy for bug reports and questions.
- Format the code and run tests: with no `main`, it runs `Test` functions.
- Several files: separate them with `-- name.go --` lines.

## Sandbox limits

- No outbound network access and no real filesystem.
- Time starts at `2009-11-10 23:00:00 UTC` and sleeps are simulated, so output is deterministic.

:::caution[Pitfalls]
- Don't paste secrets: shared snippets are public to anyone with the link.
- `time.Now()` is fake, so time-based code behaves differently than locally.
:::
