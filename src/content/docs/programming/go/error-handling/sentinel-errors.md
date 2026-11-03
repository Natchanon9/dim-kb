---
title: "Sentinel Errors"
description: "Exported error values callers compare against"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

A sentinel is a package-level error value that callers check with `errors.Is`. Name it `ErrXxx`.

## Example

```go title="stock.go"
var ErrOutOfStock = errors.New("out of stock")

func buy(stock int) error {
	if stock == 0 {
		return fmt.Errorf("buy: %w", ErrOutOfStock)
	}
	return nil
}

if err := buy(0); errors.Is(err, ErrOutOfStock) {
	fmt.Println("show sold-out badge")
}
```

Standard ones you'll meet: `io.EOF`, `sql.ErrNoRows`, `context.Canceled`, `context.DeadlineExceeded`, `os.ErrNotExist`.

## Sentinel or custom type?

Sentinel when the caller only needs to know **which** error. Custom type when it needs **data** from the error.

:::caution[Pitfalls]
- A sentinel is part of your public API; renaming or removing it breaks callers.
:::
