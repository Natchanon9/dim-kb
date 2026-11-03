---
title: "Reading Struct Tags"
description: "Build validation and mapping from tags"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

`StructField.Tag.Get("key")` reads a tag value. That is the whole mechanism behind `json`, `db` and `validate` tags.

## A tiny validator

```go title="validate.go"
type Signup struct {
	Email string `validate:"required"`
	Nick  string
}

func missingRequired(v any) []string {
	val := reflect.ValueOf(v)
	t := val.Type()
	var missing []string
	for i := range t.NumField() {
		if t.Field(i).Tag.Get("validate") == "required" && val.Field(i).IsZero() {
			missing = append(missing, t.Field(i).Name)
		}
	}
	return missing
}

missingRequired(Signup{Nick: "x"}) // [Email]
```

`Tag.Lookup(key)` also tells you whether the key exists at all.

:::caution[Pitfalls]
- In production use a maintained library like `go-playground/validator` instead of writing your own.
:::
