---
title: "sort Package"
description: "การเรียงลำดับและค้นหา ทั้งแบบเก่าและแบบใหม่"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

`sort` มีมาก่อน generics โค้ดใหม่ส่วนใหญ่ใช้ `slices.Sort`/`slices.SortFunc` แต่จะเจอ `sort` ในโค้ดทั่วไปอยู่ดี

## ตัวอย่าง

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

| แบบเก่า | แบบใหม่ (1.21+) |
| --- | --- |
| `sort.Ints`, `sort.Strings` | `slices.Sort` |
| `sort.Slice(s, less)` | `slices.SortFunc(s, cmp)` |
| `sort.SliceStable` | `slices.SortStableFunc` |
| `sort.Search` | `slices.BinarySearch(Func)` |

:::caution[ข้อควรระวัง]
- `sort.Slice` ไม่ stable: ตัวที่เท่ากันอาจสลับที่ ถ้าลำดับสำคัญให้ใช้แบบ stable
:::
