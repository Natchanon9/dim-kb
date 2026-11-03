---
title: "NULL Handling"
description: "Scan NULL columns with sql.Null types or pointers"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

Scanning a SQL `NULL` into a plain `string` or `int` fails. Use a type that can say "no value".

## Example

```go title="null.go"
var phone sql.NullString
phone.Scan(nil) // what rows.Scan does for NULL
fmt.Println(phone.Valid, phone.String == "") // false true

var age sql.Null[int64] // generic version, Go 1.22+
age.Scan(int64(30))
fmt.Println(age.Valid, age.V) // true 30

var nick *string // pointer: nil means NULL
```

| Option | Pros |
| --- | --- |
| `sql.NullString`, `sql.NullInt64`, … | Explicit, no allocation |
| `sql.Null[T]` (1.22+) | One generic type for everything |
| `*T` | Natural in JSON (`null`), simple |
| `COALESCE(col, '')` in SQL | Plain Go types when NULL = empty |

:::caution[Pitfalls]
- The error "converting NULL to string is unsupported" means a nullable column scanned into a plain type.
:::
