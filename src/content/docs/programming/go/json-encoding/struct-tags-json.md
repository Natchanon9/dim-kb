---
title: "Struct Tags for JSON"
description: "Rename, omit and reshape fields with json tags"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

The `json` tag controls how each exported field is encoded and decoded.

## Example

```go title="product.go"
type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price,omitempty"` // dropped when zero
	Stock    int     `json:"stock,string"`    // encoded as a JSON string
	Internal string  `json:"-"`               // never encoded
}

json.Marshal(Product{ID: 1, Name: "pen", Stock: 5, Internal: "x"})
// {"id":1,"name":"pen","stock":"5"}
```

| Tag | Effect |
| --- | --- |
| `json:"name"` | Key name |
| `json:",omitempty"` | Skip zero values (`0`, `""`, `nil`, empty slice/map) |
| `json:",omitzero"` | Skip zero values including zero structs and `time.Time` (Go 1.24+) |
| `json:",string"` | Number or bool as a JSON string |
| `json:"-"` | Ignore the field |

:::caution[Pitfalls]
- `omitempty` doesn't skip an empty struct or zero `time.Time`; use a pointer or `omitzero`.
- Without tags, keys are the Go field names (`"ID"`), which looks odd in APIs.
:::
