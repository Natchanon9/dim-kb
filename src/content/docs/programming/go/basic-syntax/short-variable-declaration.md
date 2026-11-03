---
title: "Short Variable Declaration"
description: "The := operator: declare and assign in one step"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

`:=` declares and initialises a variable, inferring its type. It is the default inside functions.

## Using `:=`

```go title="short.go"
count := 10
name, ok := "go", true
count, extra := 20, 1 // allowed: extra is new, count is reassigned
fmt.Println(count, name, ok, extra) // 20 go true 1
```

At least one variable on the left must be new. The others are simply reassigned, which is why `v, err := f()` works repeatedly in one function.

## `:=` vs `var`

| Use | When |
| --- | --- |
| `x := value` | Inside functions, you have a value |
| `var x T` | You want the zero value, or you are at package level |

:::caution[Pitfalls]
- `:=` inside `if`/`for` creates a **new** variable that shadows the outer one. A classic bug: `err` assigned inside a block never reaches the outer `err`.
- `:=` does not work at package level.
:::
