---
title: "Input & Output"
description: "Printing with fmt verbs and reading input with fmt.Scan"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

The `fmt` package prints and scans. Learn a handful of format verbs and you can read any debug output.

## Printing

```go title="print.go"
u := user{"Ann", 30}
fmt.Printf("%v | %+v | %T\n", u, u, u)
// {Ann 30} | {Name:Ann Age:30} | main.user
fmt.Printf("[%6.2f] [%-4s] %q\n", 3.14159, "go", "hi")
// [  3.14] [go  ] "hi"
```

| Verb | Prints |
| --- | --- |
| `%v` / `%+v` / `%#v` | Value / with field names / as Go syntax |
| `%T` | Type |
| `%d` `%f` `%s` `%q` `%t` | int, float, string, quoted string, bool |
| `%6.2f` / `%-4s` | Width and precision / left-align |

`Println` adds spaces and a newline, `Printf` formats, `Sprintf` returns a string, `Fprintf` writes to any `io.Writer`.

## Reading input

```go title="scan.go"
var name string
var age int
fmt.Scan(&name, &age) // reads "Bob 25" from stdin
```

For whole lines, use `bufio.Scanner` on `os.Stdin`.

:::caution[Pitfalls]
- `Scan` needs pointers (`&name`); passing the value does nothing useful.
- Check the error from `Scan`; bad input leaves variables unchanged.
:::
