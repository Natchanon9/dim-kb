---
title: "Inspecting Structs"
description: "Loop over struct fields with reflection"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

`Type.NumField` and `Type.Field(i)` walk a struct's fields: name, type, tag, whether exported.

## Example

```go title="fields.go"
type Account struct {
	Name  string
	Age   int
	email string
}

t := reflect.TypeOf(Account{})
for i := range t.NumField() {
	f := t.Field(i)
	fmt.Println(f.Name, f.Type, f.IsExported())
}
// Name string true
// Age int true
// email string false
```

Use `reflect.ValueOf(s).Field(i)` alongside to read the matching value. Embedded fields show up with `f.Anonymous == true`.

:::caution[Pitfalls]
- If you have a pointer, call `.Elem()` first; `NumField` on a pointer type panics.
:::
