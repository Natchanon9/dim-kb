---
title: "For Loop"
description: "for is Go's only loop: classic, while-style, infinite and range over int"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

Go has one loop keyword, `for`. It covers every loop shape other languages split into `for`, `while` and `do`.

## Four shapes

```go title="loops.go"
sum := 0
for i := 0; i < 5; i++ { // classic three-part loop
	sum += i
}
n := 1
for n < 100 { // "while" loop
	n *= 2
}
count := 0
for { // infinite loop
	count++
	if count == 3 {
		break
	}
}
for i := range 3 { // Go 1.22+: range over an int → 0 1 2
	fmt.Print(i, " ")
}
fmt.Println(sum, n, count) // 10 128 3
```

## Loop variables (Go 1.22)

Since Go 1.22 each iteration gets a **new** copy of the loop variable, so closures and goroutines started in the loop see the value from their own iteration.

:::caution[Pitfalls]
- On Go < 1.22 (check `go` in `go.mod`), closures in a loop share one variable. Old code fixes it with `i := i`.
- No parentheses around the loop header, and the braces are required.
:::
