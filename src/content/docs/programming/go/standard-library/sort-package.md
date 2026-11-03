---
title: "sort Package"
description: "Sorting and searching, old and new style"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

`sort` predates generics. New code mostly uses `slices.Sort`/`slices.SortFunc`, but you'll read `sort` everywhere.

## Example

```go title="sort.go"
ints := []int{3, 1, 2}
sort.Ints(ints)                         // [1 2 3]
sort.SearchInts(ints, 2)                // 1

people := []person{{"b", 30}, {"a", 30}, {"c", 20}}
sort.SliceStable(people, func(i, j int) bool { return people[i].Age < people[j].Age })
// [{c 20} {b 30} {a 30}]

slices.SortStableFunc(people, func(x, y person) int { return cmp.Compare(x.Name, y.Name) })
// [{a 30} {b 30} {c 20}]
```

| Old | New (1.21+) |
| --- | --- |
| `sort.Ints`, `sort.Strings` | `slices.Sort` |
| `sort.Slice(s, less)` | `slices.SortFunc(s, cmp)` |
| `sort.SliceStable` | `slices.SortStableFunc` |
| `sort.Search` | `slices.BinarySearch(Func)` |

:::caution[Pitfalls]
- `sort.Slice` is not stable: equal elements may swap. Use the stable variant when order matters.
:::
