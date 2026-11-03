---
title: "Effective Go"
description: "The core idioms that make Go code look like Go"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

"Effective Go" (go.dev/doc/effective_go) and the Go Code Review Comments wiki define the house style. A few idioms cover most of it.

## Make the zero value useful

```go title="buffer.go"
type Buffer struct{ data []byte }

func (b *Buffer) Write(p []byte) { b.data = append(b.data, p...) }

var b Buffer // no constructor needed
b.Write([]byte("go"))
```

## Core idioms

- **gofmt everything;** no style debates.
- **Handle errors right away,** return early, keep the happy path unindented.
- **Short names** for short scopes; no `Get` prefix on getters.
- **Accept interfaces, return structs;** keep interfaces small.
- **Share memory by communicating** where it makes code clearer.
- **Clear is better than clever;** a little copying beats a little dependency.

:::caution[Pitfalls]
- Porting patterns from Java/C# (deep inheritance trees, factories for everything, getters/setters) produces un-idiomatic Go.
:::
