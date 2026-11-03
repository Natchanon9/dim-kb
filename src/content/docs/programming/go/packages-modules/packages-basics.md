---
title: "Packages Basics"
description: "One folder, one package; capital letters export"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

A package is a folder of `.go` files that share one `package` name. It is Go's unit of code organisation and visibility.

## Layout

```text
myapp/
├─ go.mod            module github.com/me/myapp
├─ main.go           package main
└─ greet/
   └─ greet.go       package greet
```

```go title="greet/greet.go"
package greet

// Hello is exported: other packages can call greet.Hello.
func Hello(name string) string { return "Hello, " + name }

// suffix is unexported: only code in package greet can use it.
func suffix() string { return "!" }
```

```go title="main.go"
import "github.com/me/myapp/greet" // module path + folder

fmt.Println(greet.Hello("gopher")) // Hello, gopher
```

## Rules

- All files in a folder belong to the same package (except `_test` packages).
- The package name is usually the folder name: short, lowercase, no underscores.
- Visibility is per **package**, not per file.

:::caution[Pitfalls]
- Import cycles (`a` imports `b` imports `a`) don't compile; move shared code into a third package.
- Avoid catch-all names like `utils`, `common`, `helpers`.
:::
