---
title: "Implicit Interface Satisfaction"
description: "No implements keyword: having the methods is enough"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

Go has no `implements` keyword. A type satisfies an interface as soon as it has the methods, even if the interface was written later, in another package.

## Example

```go title="logger.go"
type Logger interface{ Log(msg string) }

type ConsoleLogger struct{ prefix string }

func (c ConsoleLogger) Log(msg string) { fmt.Println(c.prefix + msg) }

var _ Logger = ConsoleLogger{} // optional compile-time check

func start(l Logger) { l.Log("started") }
start(ConsoleLogger{"[app] "}) // [app] started

var w io.Writer = &bytes.Buffer{} // bytes.Buffer never mentions io.Writer
```

## Why it's useful

You can define a small interface for code you don't own (a third-party client) and swap in a fake for tests, without touching that code.

:::caution[Pitfalls]
- A typo in a method name means the type silently doesn't satisfy the interface; the `var _ I = T{}` line turns that into a compile error.
:::
