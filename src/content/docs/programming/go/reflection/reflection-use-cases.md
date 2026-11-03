---
title: "Reflection Use Cases"
description: "Where reflection earns its cost"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

Reflection belongs in generic infrastructure that must handle types it doesn't know at compile time.

## A mini encoder

```go title="tomap.go"
func toMap(v any) map[string]any {
	val := reflect.ValueOf(v)
	t := val.Type()
	out := map[string]any{}
	for i := range t.NumField() {
		key := t.Field(i).Tag.Get("json")
		if key == "" {
			key = t.Field(i).Name
		}
		out[key] = val.Field(i).Interface()
	}
	return out
}
// toMap(struct{ID int `json:"id"`; Name string `json:"name"`}{1, "go"}) → map[id:1 name:go]
```

## Real users

| Area | Example |
| --- | --- |
| Encoding | `encoding/json`, `encoding/xml`, YAML |
| Validation | `go-playground/validator` |
| ORMs / SQL mapping | GORM, sqlx |
| Testing | `reflect.DeepEqual`, `go-cmp` |
| Config loading | env/flag → struct libraries |

:::caution[Pitfalls]
- If you're writing reflection in business logic, an interface or generics is almost always the better answer.
:::
