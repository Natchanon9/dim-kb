---
title: "iota"
description: "constant ที่นับเพิ่มเอง ใช้ทำ enum และขนาดแบบ bit"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

`iota` เป็นตัวนับในบล็อก `const (...)` บรรทัดแรกเป็น 0 แล้วเพิ่มทีละ 1 ต่อบรรทัด Go ใช้แทน keyword `enum`

## Enum

```go title="iota.go"
type Status int

const (
	Pending Status = iota // 0
	Active                // 1
	Closed                // 2
)
```

บรรทัดที่ไม่ใส่ค่าจะใช้ expression เดิมซ้ำกับ `iota` ตัวถัดไป

## Expression และการข้ามค่า

```go title="sizes.go"
const (
	_  = iota             // ข้าม 0
	KB = 1 << (10 * iota) // 1 << 10
	MB                    // 1 << 20
)
fmt.Println(Pending, Active, Closed, KB, MB) // 0 1 2 1024 1048576
```

:::caution[ข้อควรระวัง]
- `iota` กลับเป็น 0 ทุกครั้งที่เริ่มบล็อก `const` ใหม่
- ถ้า 0 หมายถึง "ยังไม่ตั้งค่า" ให้เริ่มค่าจริงที่ `iota + 1` หรือข้าม 0 ด้วย `_` จะได้ไม่สับสน field ที่เป็น zero value กับสถานะจริง
:::
