---
title: "slices Package"
description: "sort ค้นหา และเปรียบเทียบ slice ด้วย standard library (Go 1.21+)"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

package `slices` แบบ generic (Go 1.21+) ใช้แทน loop ที่เขียนเองกับ slice ได้เกือบหมด

## Function ที่ใช้บ่อย

```go title="pkg.go"
s := []int{3, 1, 2}
slices.Sort(s)
i, found := slices.BinarySearch(s, 2)
fmt.Println(s, slices.Contains(s, 3), slices.Index(s, 3), i, found, slices.Max(s))
// [1 2 3] true 2 1 true 3
```

| ต้องการ | Function |
| --- | --- |
| Sort | `Sort`, `SortFunc`, `SortStableFunc` |
| ค้นหา | `Contains`, `Index`, `IndexFunc`, `BinarySearch` |
| เปรียบเทียบ | `Equal`, `Compare` |
| แก้ไข | `Insert`, `Delete`, `DeleteFunc`, `Compact`, `Reverse` |
| Copy | `Clone` |
| ค่าต่ำสุด / สูงสุด | `Min`, `Max` |

`SortFunc` รับ comparator ที่คืนค่าลบ/ศูนย์/บวก ส่วนใหญ่ใช้ `cmp.Compare(a.Age, b.Age)`

:::caution[ข้อควรระวัง]
- `BinarySearch` ใช้ได้กับ slice ที่ sort แล้วเท่านั้น
- `Min`/`Max` panic ถ้า slice ว่าง
:::
