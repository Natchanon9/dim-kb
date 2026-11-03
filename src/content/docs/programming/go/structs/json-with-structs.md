---
title: "JSON with Structs"
description: "Marshal and unmarshal structs with encoding/json"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

`encoding/json` maps exported struct fields to JSON keys, using tags to rename or skip them.

## Encode and decode

```go title="order.go"
type Order struct {
	ID    int      `json:"id"`
	Items []string `json:"items"`
	Note  string   `json:"note,omitempty"`
}

data, _ := json.Marshal(Order{ID: 1, Items: []string{"pen"}})
// {"id":1,"items":["pen"]}

var o Order
err := json.Unmarshal([]byte(`{"id":2,"items":["a","b"],"extra":true}`), &o)
// o.ID=2, o.Items=[a b], err=nil (unknown "extra" is ignored)
```

Key matching on decode is case-insensitive; fields missing from the JSON keep their zero values.

:::caution[Pitfalls]
- `Unmarshal` needs a pointer (`&o`).
- You can't tell "missing" from "zero" with plain fields; use a pointer field (`*int`) when that matters.
- Use `json.NewDecoder(r.Body)` with `DisallowUnknownFields()` to reject unexpected fields in APIs.
:::
