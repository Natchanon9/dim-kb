---
title: "Exported & Unexported Fields"
description: "Capitalised fields are public; lowercase fields are package-private"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

A field starting with a capital letter is exported: other packages can read and write it. Lowercase fields are visible only inside the declaring package.

## Effect on encoding

```go title="account.go"
type account struct {
	ID       int    // exported: visible to other packages and to encoding/json
	password string // unexported: package-private
}

b, _ := json.Marshal(account{ID: 1, password: "x"})
fmt.Println(string(b)) // {"ID":1}
```

`encoding/json`, `encoding/xml` and ORMs live in other packages, so they only see exported fields.

## Encapsulation

Keep fields unexported when you need to protect invariants, and expose methods instead (`func (a *account) SetPassword(p string) error`).

:::caution[Pitfalls]
- "My JSON is `{}`" almost always means the fields are lowercase.
:::
