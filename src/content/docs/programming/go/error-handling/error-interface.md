---
title: "error Interface"
description: "Errors are ordinary values implementing Error() string"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

`error` is a built-in interface with one method. Any type with `Error() string` is an error, and `nil` means "no error".

## The interface

```go title="error.go"
// built in:
// type error interface { Error() string }

var err error = errors.New("disk full")
fmt.Println(err, err.Error() == "disk full") // disk full true
```

## Errors are values

Go has no exceptions. A function that can fail returns an `error` as its last result, and the caller checks it right away:

```go
f, err := os.Open("config.yaml")
if err != nil {
	return err
}
```

Because errors are values, you can store them, compare them, wrap them and return them like anything else.

:::caution[Pitfalls]
- Don't compare error **messages** (`err.Error() == "..."`); messages change. Use `errors.Is`/`errors.As`.
:::
