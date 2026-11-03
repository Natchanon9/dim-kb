---
title: "Go Workspaces"
description: "Develop several local modules together with go.work"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

A workspace (Go 1.18+) lets you build several modules on your disk as if they were one, without `replace` lines in `go.mod`.

## Setup

```sh
go work init ./api ./shared   # creates go.work
go work use ./worker          # add another module
```

```text title="go.work"
go 1.22

use (
	./api
	./shared
	./worker
)
```

Now `api` imports the local `shared` folder directly, and edits show up immediately.

## When to use

Changing a library and the service that uses it at the same time, or a monorepo with several modules.

:::caution[Pitfalls]
- Many teams keep `go.work` out of git so CI builds each module against its real `go.mod`. Set `GOWORK=off` to test without it.
:::
