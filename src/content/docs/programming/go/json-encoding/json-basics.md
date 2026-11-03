---
title: "JSON Basics"
description: "Marshal and Unmarshal with encoding/json"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

`json.Marshal` turns Go values into JSON bytes; `json.Unmarshal` fills a Go value from JSON. Structs, maps, slices and basic types all work.

## Example

```go title="json.go"
b, _ := json.Marshal(map[string]any{"name": "go", "tags": []string{"a"}, "ok": true})
// {"name":"go","ok":true,"tags":["a"]}  (map keys sorted)

var v map[string]any
json.Unmarshal([]byte(`{"n":1}`), &v)
fmt.Printf("%T\n", v["n"]) // float64

pretty, _ := json.MarshalIndent(map[string]int{"a": 1}, "", "  ")
```

## Type mapping

| JSON | Go (into `any`) | Go (typed) |
| --- | --- | --- |
| object | `map[string]any` | struct, `map[string]T` |
| array | `[]any` | slice, array |
| number | `float64` | any int/float type |
| string | `string` | `string`, `time.Time`, `[]byte` (base64) |
| `true`/`false` | `bool` | `bool` |
| `null` | `nil` | pointer, slice, map → nil |

Prefer decoding into structs over `map[string]any`.

:::caution[Pitfalls]
- Large integer IDs lose precision as `float64`; decode into `int64` fields or use `dec.UseNumber()`.
:::
