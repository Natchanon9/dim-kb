---
title: "Documentation"
description: "Doc comments, examples and pkg.go.dev"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

Go documentation is just comments. A comment right above a declaration becomes its docs in `go doc`, editors and pkg.go.dev.

## Doc comments

```go title="greet.go"
// Greet returns a greeting for name.
//
// It is safe for concurrent use. See also [Greeter.Greet].
func Greet(name string) string { return "Hello, " + name }
```

- Start with the name being documented, as a full sentence.
- `[Name]` links to another identifier (Go 1.19+).
- A `// Package foo ...` comment in one file (often `doc.go`) documents the package.

## Examples are tested docs

```go title="greet_test.go"
func ExampleGreet() {
	fmt.Println(Greet("docs"))
	// Output: Hello, docs
}
```

`go test` runs it and checks the output; pkg.go.dev shows it under `Greet`.

:::caution[Pitfalls]
- A comment separated from the declaration by a blank line is not a doc comment.
:::
