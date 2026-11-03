---
title: "Internal Packages"
description: "Hide packages from code outside your module with internal/"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

A package under a folder named `internal` can only be imported by code rooted at the parent of that `internal` folder. The compiler enforces it.

## Rule

```text
myapp/
├─ internal/format/    importable by myapp/... only
├─ greet/              can import internal/format
└─ cmd/server/         can import internal/format
```

```go title="greet/greet.go"
import "github.com/me/myapp/internal/format"

func Hello(name string) string { return format.Title("hello, " + name) } // Hello, ...
```

Another module importing `github.com/me/myapp/internal/format` gets a compile error: "use of internal package not allowed".

## Why use it

Keep implementation details private so you can change them without breaking anyone. In services, most code lives in `internal/`.

:::caution[Pitfalls]
- `internal` restricts importers, not visibility of names; exported names inside are still exported to allowed importers.
:::
