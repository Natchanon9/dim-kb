---
title: "Table-Driven Tests"
description: "loop เดียวบน slice ของ test case"
sidebar:
  order: 2
status: done
lastReviewed: 2026-11-03
---

ใส่ input และค่าที่คาดไว้ใน slice ของ struct แล้ววน loop เพิ่ม case ใหม่แค่บรรทัดเดียว

## ตัวอย่าง

```go title="abs_test.go"
func TestAbs(t *testing.T) {
	tests := []struct {
		name     string
		in, want int
	}{
		{"positive", 3, 3},
		{"negative", -3, 3},
		{"zero", 0, 0},
	}
	for _, tt := range tests {
		if got := Abs(tt.in); got != tt.want {
			t.Errorf("%s: Abs(%d) = %d, want %d", tt.name, tt.in, got, tt.want)
		}
	}
}
```

ใช้ร่วมกับ `t.Run` (หน้าถัดไป) ให้แต่ละ case เป็น subtest ที่มีชื่อและรันแยกได้

:::caution[ข้อควรระวัง]
- ใส่ field `name` ด้วย "case 3 failed" หายาก
- ใน loop ใช้ `t.Errorf` ไม่ใช่ `t.Fatalf` จะได้ไม่มี case ที่ fail ตัวเดียวบังตัวอื่น
:::
