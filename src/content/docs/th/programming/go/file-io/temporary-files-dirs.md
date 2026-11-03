---
title: "ไฟล์และโฟลเดอร์ชั่วคราว"
description: "สร้างไฟล์ชั่วคราวอย่างปลอดภัยและลบทิ้งเมื่อเสร็จ"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

`os.CreateTemp` และ `os.MkdirTemp` สร้างไฟล์และโฟลเดอร์ชื่อไม่ซ้ำในโฟลเดอร์ temp ของระบบ เราต้องลบเอง

## ตัวอย่าง

```go title="temp.go"
f, err := os.CreateTemp("", "upload-*.txt") // * → string สุ่ม
if err != nil {
	return err
}
defer os.Remove(f.Name())
f.WriteString("tmp")
f.Close()
// ชื่อจะประมาณ /tmp/upload-123456789.txt
```

## ใน test

```go
dir := t.TempDir() // ลบให้อัตโนมัติเมื่อ test จบ
```

:::caution[ข้อควรระวัง]
- `defer os.Remove`/`os.RemoveAll` เสมอ โฟลเดอร์ temp บน server เต็มแบบเงียบๆ
- อย่าตั้งชื่อไฟล์ temp ตายตัวเอง รันพร้อมกันแล้วชนกัน
:::
