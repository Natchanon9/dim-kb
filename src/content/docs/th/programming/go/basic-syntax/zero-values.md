---
title: "Zero Values"
description: "ตัวแปรทุกตัวเริ่มด้วยค่าเริ่มต้นที่ใช้ได้เลย"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

ตัวแปรที่ประกาศแต่ไม่กำหนดค่า จะได้ **zero value** ของ type นั้น Go ไม่มีสถานะ "ยังไม่ได้ initialize"

## Zero value ของแต่ละ type

| Type | Zero value |
| --- | --- |
| ตัวเลข | `0` |
| `bool` | `false` |
| `string` | `""` |
| pointer, slice, map, channel, func, interface | `nil` |
| struct | ทุก field เป็น zero value ของตัวเอง |

```go title="zero.go"
var i int
var s string
var ok bool
var p *int
var list []int
var m map[string]int
fmt.Printf("%d %q %t %v %t %t\n", i, s, ok, p, list == nil, m == nil)
// 0 "" false <nil> true true
```

## ทำให้ zero value ใช้งานได้

type ที่ออกแบบดีใช้ได้เลยโดยไม่ต้องมี constructor เช่น `sync.Mutex`, `bytes.Buffer` และ slice ที่เป็น nil (`append` ได้เลย)

:::caution[ข้อควรระวัง]
- เขียนลง map ที่เป็น nil จะ panic ต้องสร้างด้วย `make(map[K]V)` ก่อน
- เรียก method ผ่าน pointer ที่เป็น nil จะ panic ถ้า method นั้นอ่าน field
:::
