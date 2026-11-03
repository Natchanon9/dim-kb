---
title: "Struct Embedding"
description: "Compose types by embedding; fields and methods are promoted"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

Embedding puts a type inside a struct without a field name. Its fields and methods are **promoted**, so they can be used as if they belonged to the outer struct. This is composition, not inheritance.

## Example

```go title="embed.go"
type Base struct{ ID int }

func (b Base) Describe() string { return fmt.Sprintf("id=%d", b.ID) }

type Product struct {
	Base // embedded
	Name string
}

p := Product{Base{7}, "pen"}
fmt.Println(p.ID, p.Describe(), p.Base.ID) // 7 id=7 7
```

## Not inheritance

- `Product` is not a `Base`; you can't pass it where a `Base` is expected.
- Promoted methods still receive the inner `Base`, not the `Product`.
- The outer type can **shadow** a promoted field or method by defining its own.

Embedding interfaces in interfaces, and `sync.Mutex` in structs, are common real uses.

:::caution[Pitfalls]
- Embedding a type exports its whole method set from your type; embed only what you want public.
- Two embedded types with the same field name make that name ambiguous; access it through the type name.
:::
