---
title: "maps Package"
description: "clone เปรียบเทียบ copy และกรอง map (Go 1.21+)"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

package `maps` แบบ generic (Go 1.21+) แทน loop ที่ต้องเขียนเองกับ map ได้

## Function ที่ใช้บ่อย

```go title="pkg.go"
a := map[string]int{"x": 1}
b := maps.Clone(a)
b["y"] = 2
fmt.Println(maps.Equal(a, b), len(a), len(b)) // false 1 2

maps.Copy(a, b) // copy entry ของ b ลง a
maps.DeleteFunc(a, func(k string, v int) bool { return v > 1 })
fmt.Println(a) // map[x:1]
```

| Function | ทำอะไร |
| --- | --- |
| `Clone` | copy แบบตื้น |
| `Equal`, `EqualFunc` | เปรียบเทียบเนื้อหา |
| `Copy(dst, src)` | ใส่/เขียนทับทุก entry ของ src ลง dst |
| `DeleteFunc` | ลบ entry ที่ตรงเงื่อนไข |
| `Keys`, `Values` (1.23+) | iterator ของ key / value |

:::caution[ข้อควรระวัง]
- `Clone` เป็นแบบตื้น value ที่เป็น slice หรือ pointer จะแชร์กันระหว่าง map ทั้งสอง
:::
