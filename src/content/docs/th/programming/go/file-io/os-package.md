---
title: "os Package"
description: "อ่าน เขียน เปิด และตรวจไฟล์"
sidebar:
  order: 1
status: done
lastReviewed: 2026-11-03
---

`os` คือประตูระดับล่างสู่ filesystem มี helper แบบอ่านเขียนทั้งไฟล์สำหรับไฟล์เล็ก และ `*os.File` สำหรับ stream

## Helper แบบทั้งไฟล์

```go title="files.go"
path := filepath.Join(dir, "note.txt")
err := os.WriteFile(path, []byte("hello\n"), 0o644)
data, err := os.ReadFile(path) // "hello\n"

_, err = os.Stat(filepath.Join(dir, "missing"))
errors.Is(err, os.ErrNotExist) // true
```

## เปิดไฟล์

| เรียก | ใช้เมื่อ |
| --- | --- |
| `os.Open(p)` | อ่านอย่างเดียว |
| `os.Create(p)` | สร้างหรือล้างไฟล์ แล้วเขียน |
| `os.OpenFile(p, os.O_APPEND\|os.O_CREATE\|os.O_WRONLY, 0o644)` | ต่อท้าย หรือกำหนด flag เอง |
| `os.MkdirAll(p, 0o755)` | สร้างโฟลเดอร์ |
| `os.Remove` / `os.RemoveAll` | ลบ |

```go
f, err := os.Open(path)
if err != nil {
	return err
}
defer f.Close()
```

:::caution[ข้อควรระวัง]
- เช็ค error ของ `Close` กับไฟล์ที่ **เขียน** เพราะอาจบอกว่า flush ไม่สำเร็จ
- `ReadFile` โหลดทั้งก้อนเข้า memory ไฟล์ใหญ่ให้ stream ด้วย `bufio`/`io.Copy`
:::
