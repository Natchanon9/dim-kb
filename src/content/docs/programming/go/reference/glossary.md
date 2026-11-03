---
title: "Glossary"
description: "Terms used on this site, kept consistent across English / Thai"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

These terms stay in English in the Thai pages. The meaning column is the one-line definition used across the site.

| Term | Meaning |
| --- | --- |
| package | Folder of `.go` files sharing one name; unit of visibility |
| module | Tree of packages with a `go.mod`; unit of versioning |
| exported | Name starting with a capital letter, visible to other packages |
| zero value | Default value of a type when not set (`0`, `""`, `nil`) |
| slice | Growable view over an array: pointer, length, capacity |
| backing array | The array a slice points into; may be shared |
| map | Hash table from comparable keys to values |
| rune | One Unicode code point (`int32`) |
| struct | Named group of fields |
| embedding | Putting a type in a struct without a field name; promotes its members |
| method | Function with a receiver |
| receiver | The value a method is called on (value or pointer) |
| interface | Set of method signatures, satisfied implicitly |
| type assertion | `x.(T)`: get the concrete value out of an interface |
| generics / type parameter | `[T any]`: code written once for many types |
| constraint | Interface limiting which types a type parameter accepts |
| pointer | Address of a value; `&x` takes it, `*p` follows it |
| goroutine | Lightweight concurrent function managed by the Go runtime |
| channel | Typed pipe for sending values between goroutines |
| data race | Unsynchronised concurrent access with at least one write |
| deadlock | Every goroutine is blocked waiting on another |
| context | Carries cancellation, deadline and request values |
| middleware | Wrapper around a handler that runs before/after it |
| handler | Function or type that serves an HTTP request |
| sentinel error | Exported error value compared with `errors.Is` |
| wrap | Add context to an error while keeping the original (`%w`) |
| panic / recover | Abort the goroutine / catch that abort in a deferred call |
| build tag | `//go:build` constraint deciding which files compile |
| port / adapter | Interface the core needs / its concrete implementation |

:::caution[Pitfalls]
- When a Thai page needs a new term, add it here first so both languages use the same word.
:::
