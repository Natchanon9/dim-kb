---
title: "Project Layout"
description: "cmd/, internal/ and keeping structure flat"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

Go has no mandatory layout. Start flat and add folders when the code asks for them. A common shape for services:

```text
myapp/
├─ go.mod
├─ cmd/
│  ├─ api/main.go        one folder per binary
│  └─ worker/main.go
├─ internal/             private to this module
│  ├─ order/             a feature: handler, service, repository
│  ├─ user/
│  └─ platform/          db, config, logging setup
├─ migrations/           SQL files
└─ Dockerfile
```

## Guidelines

- **`main` stays thin:** read config, wire dependencies, start the server.
- **Organise by feature** (`order`, `user`) rather than by layer (`controllers`, `models`).
- **`internal/`** for everything not meant to be imported by other modules.
- **`pkg/`** only if you really publish reusable packages; many projects don't need it.
- **Small project?** `main.go` plus a few files in the root is fine.

:::caution[Pitfalls]
- "golang-standards/project-layout" is a community repo, not an official standard; don't copy every folder.
:::
