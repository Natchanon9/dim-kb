---
title: "Importing Packages"
description: "Import paths, grouping, aliases and blank imports"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

An import path is the module path plus the folder. Inside the file you refer to the package by its name.

## Forms

```go title="imports.go"
import (
	"fmt"
	str "strings" // alias

	g "github.com/me/myapp/greet"
	_ "github.com/lib/pq" // side effects only (registers a driver)
)

fmt.Println(str.Repeat("go", 2), g.Hello("alias")) // gogo Hello, alias
```

| Form | Use |
| --- | --- |
| `import "pkg"` | Normal |
| `import alias "pkg"` | Name clashes (`crand "crypto/rand"`) or long names |
| `import _ "pkg"` | Run the package's `init` only |
| `import . "pkg"` | Avoid; hides where names come from |

`goimports` (or your editor) adds, removes and groups imports: standard library first, then third-party.

:::caution[Pitfalls]
- An unused import is a compile error.
:::
