---
title: "Struct Tags"
description: "Metadata strings on fields for json, db and validation libraries"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

A tag is a string attached to a field. Packages read it with reflection to decide how to encode, store or validate that field.

## Example

```go title="tags.go"
type Item struct {
	Name   string `json:"name"`
	Price  int    `json:"price,omitempty"`
	Secret string `json:"-"`
}

b, _ := json.Marshal(Item{Name: "pen", Secret: "x"})
f, _ := reflect.TypeOf(Item{}).FieldByName("Name")
fmt.Println(string(b), f.Tag.Get("json")) // {"name":"pen"} name
```

## Common tags

| Tag | Used by |
| --- | --- |
| `json:"name,omitempty"` | `encoding/json` |
| `db:"user_id"` | sqlx and similar |
| `gorm:"primaryKey"` | GORM |
| `validate:"required,email"` | go-playground/validator |

Several tags go in one string, separated by spaces: `` `json:"id" db:"id"` ``.

:::caution[Pitfalls]
- The format is strict: `key:"value"`, no spaces around the colon. A typo silently does nothing; `go vet` catches malformed tags.
:::
