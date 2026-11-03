---
title: "การ Format String"
description: "verb ของ Sprintf สำหรับตัวเลข ความกว้าง การเติม และการจัดชิด"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

`fmt.Sprintf` สร้าง string ตาม format ด้วย verb ชุดเดียวกับ `Printf` ความกว้างและ flag ใส่ระหว่าง `%` กับ verb

## Verb และ flag

```go title="format.go"
fmt.Sprintf("%05d|%x|%X|%o|%b", 42, 255, 255, 8, 5)
// 00042|ff|FF|10|101
fmt.Sprintf("|%8.3f|%-8s|%8s|", 3.14159, "left", "right")
// |   3.142|left    |   right|
fmt.Printf("%v %d%%\n", []int{1, 2}, 50)
// [1 2] 50%
```

| Flag | ความหมาย |
| --- | --- |
| `%8s` | กว้าง 8 ชิดขวา |
| `%-8s` | ชิดซ้าย |
| `%05d` | เติมศูนย์ข้างหน้า |
| `%.3f` | ทศนิยม 3 ตำแหน่ง |
| `%x` / `%X` | เลขฐาน 16 (ใช้กับ string และ `[]byte` ได้ด้วย) |
| `%%` | ตัว `%` |

type ที่มี method `String() string` จะพิมพ์ผ่าน method นั้นเมื่อใช้ `%v` และ `%s`

:::caution[ข้อควรระวัง]
- verb กับ argument ไม่ตรงกันจะพิมพ์ `%!d(string=hi)` ตอน runtime ส่วนใหญ่ `go vet` จับได้
- เรียก `fmt.Sprintf("%v", x)` ข้างใน `x.String()` จะวนเรียกตัวเองไม่จบ
:::
