---
title: "Subtests"
description: "case ย่อยที่มีชื่อด้วย t.Run การกรอง และการรันขนาน"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

`t.Run(name, func(t *testing.T) {...})` สร้าง subtest ที่มีชื่อ สถานะผ่าน/ไม่ผ่าน และขอบเขตของ `Fatal` เป็นของตัวเอง

## ตัวอย่าง

```go title="sub_test.go"
for _, tt := range tests {
	t.Run(tt.name, func(t *testing.T) {
		t.Parallel() // รันแต่ละ case พร้อมกัน
		if got := Abs(tt.in); got != tt.want {
			t.Errorf("Abs(%d) = %d, want %d", tt.in, got, tt.want)
		}
	})
}
```

```sh
go test -run 'TestAbsSubtests/negative' -v
```

## ข้อดี

- รันหรือ debug case เดียวตามชื่อได้
- `t.Fatal` หยุดแค่ subtest นั้น
- setup ที่ใช้ร่วมอยู่ก่อน loop ส่วน setup ราย case อยู่ข้างใน

:::caution[ข้อควรระวัง]
- ถ้าใช้ `t.Parallel()` กับ Go < 1.22 ต้อง `tt := tt` ก่อน ตั้งแต่ 1.22 loop variable แยกต่อรอบแล้ว
:::
