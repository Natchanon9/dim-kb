---
title: "พื้นฐาน Testing"
description: "เขียนและรัน test ด้วย go test"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

test อยู่ในไฟล์ `_test.go` ข้างๆ โค้ด test คือ function `TestXxx(t *testing.T)` แล้ว `go test` จะหาและรันให้ ไม่ต้องใช้ framework

## ตัวอย่าง

```go title="add_test.go"
func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Errorf("Add(2, 3) = %d, want 5", got)
	}
}
```

```sh
go test ./...           # ทุก package
go test -run TestAdd -v # test เดียว แบบละเอียด
go test -race ./...     # เปิด race detector
go test -count=1 ./...  # ไม่ใช้ cache ของ test
```

## การรายงานผล

| เรียก | ผล |
| --- | --- |
| `t.Errorf` / `t.Error` | บอกว่า fail แล้วทำต่อ |
| `t.Fatalf` / `t.Fatal` | บอกว่า fail แล้วหยุด test นี้ |
| `t.Skip` | ข้าม (เช่น ต้องใช้ database) |
| `t.Log` | พิมพ์เฉพาะตอน fail หรือใช้ `-v` |

ธรรมเนียมของข้อความ: `Func(args) = got, want want`

:::caution[ข้อควรระวัง]
- `t.Fatal` จาก goroutine ที่เราสร้างไม่ได้หยุด test ให้ส่งผลกลับไปที่ goroutine ของ test แทน
:::
