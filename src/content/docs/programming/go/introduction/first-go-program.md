---
title: "First Go Program"
description: "Hello World: package main, func main, go run and go build"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

Every Go program starts in `package main`, function `main`. Create a module first so the tools know your project root.

## Create and run

```sh
mkdir hello && cd hello
go mod init example.com/hello
```

```go title="main.go"
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```

```sh
go run .        # compile + run, nothing saved
go build        # writes ./hello binary
```

## What each line does

- **`package main`:** marks this as an executable, not a library.
- **`import "fmt"`:** brings in the formatting package; unused imports are a compile error.
- **`func main()`:** the entry point. No arguments, no return value; use `os.Args` and `os.Exit` instead.

:::caution[Pitfalls]
- Unused variables and imports don't compile. Delete them or use `_`.
- The opening `{` must be on the same line as `func main()`; Go inserts semicolons at line ends.
:::
