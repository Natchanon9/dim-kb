---
title: "Test Helpers"
description: "t.Helper, t.Cleanup, t.TempDir และ t.Setenv"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

helper ลดโค้ดซ้ำใน test และ `testing.T` ก็มี helper สำหรับ setup และ teardown ที่ใช้บ่อยมาให้

## ตัวอย่าง

```go title="helpers_test.go"
func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper() // ตอน fail จะชี้บรรทัดของผู้เรียก
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestWithHelpers(t *testing.T) {
	t.Setenv("APP_ENV", "test") // คืนค่าเดิมทีหลัง
	dir := t.TempDir()          // ลบทีหลัง
	t.Cleanup(func() { /* ปิด DB, หยุด server */ })

	assertEqual(t, Greet("Ann"), "Hello, Ann!")
}
```

| Helper | ทำอะไร |
| --- | --- |
| `t.Helper()` | ซ่อน function นี้จากเลขบรรทัดตอน fail |
| `t.Cleanup(f)` | รัน f เมื่อ test จบ (LIFO) |
| `t.TempDir()` | โฟลเดอร์ใหม่ ลบให้อัตโนมัติ |
| `t.Setenv(k, v)` | ตั้ง env แล้วคืนค่าเดิมทีหลัง |

ถ้าต้องการ assertion ที่ครบกว่านี้ หลายทีมใช้ `github.com/stretchr/testify` (`require`, `assert`)

:::caution[ข้อควรระวัง]
- ใช้ `t.Setenv` คู่กับ `t.Parallel()` ไม่ได้
:::
