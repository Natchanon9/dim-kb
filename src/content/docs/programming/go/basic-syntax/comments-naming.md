---
title: "Comments & Naming"
description: "Doc comments, exported names and Go naming style"
sidebar:
  order: 10
status: done
lastReviewed: 2026-11-03
---

Capitalisation controls visibility, and comments above declarations become documentation. Naming style is part of the language culture.

## Exported vs unexported

```go title="naming.go"
// MaxItems is exported: it starts with a capital letter.
const MaxItems = 10

// discount is unexported: only code in this package can call it.
func discount(price float64) float64 { return price * 0.9 }
```

## Naming rules of thumb

- **camelCase / PascalCase**, never snake_case.
- **Acronyms keep one case:** `userID`, `HTTPServer`, not `userId`.
- **Short names for short scopes:** `i`, `r`, `ctx`; longer names for package-level things.
- **No stutter:** `http.Server`, not `http.HTTPServer`.
- **Getters drop `Get`:** `u.Name()`, not `u.GetName()`.

## Doc comments

A comment directly above an exported name starts with that name and is shown by `go doc` and pkg.go.dev.

:::caution[Pitfalls]
- `//` is idiomatic; `/* */` is mostly for package docs or disabling code.
- Lint tools flag exported names without a doc comment.
:::
