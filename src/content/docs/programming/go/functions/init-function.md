---
title: "init Function"
description: "Package setup that runs before main"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

`func init()` runs automatically once per package, after package variables are set and before `main`. You can't call it yourself.

## Example

```go title="config.go"
var config map[string]string

func init() {
	config = map[string]string{"env": "dev"}
}
// config["env"] → dev
```

## Order

1. Imported packages initialise first.
2. Then package-level variables, then every `init` in the package (a file can have several).
3. Then `main`.

:::caution[Pitfalls]
- Keep `init` small. Hidden side effects (opening DBs, reading env) make code hard to test; do that in `main` and pass things in.
- `init` can't return an error; failing there means `panic` or `log.Fatal`.
:::
