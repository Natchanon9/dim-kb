---
title: "Vendoring"
description: "Copy dependencies into the repo with go mod vendor"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

`go mod vendor` copies every dependency into a `vendor/` folder. When `vendor/` exists and is consistent with `go.mod`, the go command builds from it instead of the module cache.

## Commands

```sh
go mod vendor        # create or refresh vendor/
go build ./...       # uses vendor/ automatically (Go 1.14+)
go build -mod=mod    # ignore vendor/ and use the module cache
```

## When to vendor

- Builds must work without network access (air-gapped CI).
- You want dependency code visible in code review.

Most projects don't vendor; the module proxy and `go.sum` already make builds reproducible.

:::caution[Pitfalls]
- After changing `go.mod`, rerun `go mod vendor` or builds fail with "inconsistent vendoring".
:::
