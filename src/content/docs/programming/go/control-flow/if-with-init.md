---
title: "If with Init"
description: "Run a short statement before the condition and keep its variables scoped"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

`if` can start with a statement, separated from the condition by `;`. Variables declared there exist only inside the `if`/`else` blocks.

## Pattern

```go title="init.go"
stock := map[string]int{"apple": 3}
if n, ok := stock["apple"]; ok {
	fmt.Println("in stock:", n) // in stock: 3
}
if _, err := strconv.Atoi("x"); err != nil {
	fmt.Println("error:", err)
}
// n and err don't exist here
```

Use it for comma-ok checks and for errors you handle immediately. It keeps short-lived variables out of the rest of the function.

:::caution[Pitfalls]
- If you need the value after the `if`, declare it before; otherwise it's gone.
- Keep the init statement short. A long call in there is hard to read.
:::
