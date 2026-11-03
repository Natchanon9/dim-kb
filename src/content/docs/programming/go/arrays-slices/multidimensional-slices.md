---
title: "Multidimensional Slices"
description: "Slices of slices for grids and matrices"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

Go has no built-in 2D slice. A grid is a slice of rows, and every row is a separate slice you must create.

## Building a grid

```go title="grid.go"
grid := make([][]int, 2)
for i := range grid {
	grid[i] = make([]int, 3) // each row is its own slice
}
grid[1][2] = 9
fmt.Println(grid) // [[0 0 0] [0 0 9]]
```

Rows can have different lengths ("jagged"), which suits things like adjacency lists.

## Fixed size

If both dimensions are fixed, an array of arrays is simpler: `var board [8][8]int`.

:::caution[Pitfalls]
- `make([][]int, 2)` gives two **nil** rows; indexing `grid[0][0]` panics until each row is made.
:::
