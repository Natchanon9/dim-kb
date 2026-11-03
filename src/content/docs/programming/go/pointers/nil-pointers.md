---
title: "Nil Pointers"
description: "Avoid nil dereference panics and write nil-safe methods"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

A nil pointer points nowhere. Reading a field through it panics, but calling a method on it is allowed, so methods can handle nil themselves.

## Nil-safe method

```go title="node.go"
type Node struct {
	Val  int
	Next *Node
}

func (n *Node) Sum() int {
	if n == nil { // a nil receiver is legal; guard it
		return 0
	}
	return n.Val + n.Next.Sum()
}

list := &Node{1, &Node{2, nil}}
var empty *Node
fmt.Println(list.Sum(), empty.Sum()) // 3 0
```

## Where nil pointers come from

- Uninitialised pointer fields in structs.
- Functions returning `nil, err` where the caller ignores `err`.
- Map lookups of `map[K]*V` for missing keys.

:::caution[Pitfalls]
- Always check `err` before using the other result; that is where most nil dereferences start.
:::
