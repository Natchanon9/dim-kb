---
title: "Pointer Receivers"
description: "Methods that modify their receiver"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

A method with a `*T` receiver works on the original value, so changes stick. Go inserts `&` or `*` for you when calling.

## Example

```go title="account.go"
type Account struct{ Balance int }

func (a *Account) Deposit(n int) { a.Balance += n }

acc := Account{}
acc.Deposit(50) // Go passes &acc automatically
ptr := &acc
ptr.Deposit(25)
fmt.Println(acc.Balance) // 75
```

Automatic `&` only works on addressable values: a variable, a slice element, a struct field. It doesn't work on a map element or a function's return value (`getAccount().Deposit(1)` fails if `getAccount` returns a value).

:::caution[Pitfalls]
- Mixing value and pointer receivers on one type is confusing; the comparison is on [Value vs Pointer Receivers](/dim-kb/programming/go/methods-interfaces/value-vs-pointer-receivers/).
:::
