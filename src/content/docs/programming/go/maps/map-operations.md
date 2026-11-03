---
title: "Map Operations"
description: "Insert, update, delete, len and clear"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

Map operations are built into the language, not methods.

## The operations

```go title="ops.go"
m := map[string]int{"a": 1, "b": 2}
m["a"] = 10       // insert or update
delete(m, "b")    // remove
delete(m, "nope") // no-op, no error
fmt.Println(m, len(m)) // map[a:10] 1
clear(m)               // Go 1.21+: remove everything
```

| Operation | Syntax |
| --- | --- |
| Insert / update | `m[k] = v` |
| Read | `v := m[k]` or `v, ok := m[k]` |
| Delete | `delete(m, k)` |
| Count | `len(m)` |
| Empty | `clear(m)` |

`fmt` prints maps with keys sorted, which makes test output stable.

:::caution[Pitfalls]
- You can't take the address of an element: `&m[k]` doesn't compile. For `map[string]Point`, `m[k].X = 1` fails too; store pointers or reassign the whole value.
:::
