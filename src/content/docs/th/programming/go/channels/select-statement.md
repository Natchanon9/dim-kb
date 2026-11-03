---
title: "Select"
description: "รอ operation ของหลาย channel พร้อมกัน"
sidebar:
  order: 6
status: done
lastReviewed: 2026-11-03
---

`select` รอจนกว่าจะมี case ใดทำงานต่อได้ แล้วรัน case นั้น ถ้าพร้อมหลาย case จะสุ่มเลือก

## Timeout และไม่บล็อก

```go title="select.go"
fast := make(chan string, 1)
fast <- "cache"
select {
case v := <-fast:
	fmt.Println(v) // cache
case <-time.After(time.Second):
	fmt.Println("timeout")
}

empty := make(chan int)
select {
case v := <-empty:
	fmt.Println(v)
default: // รันทันทีถ้าไม่มี case ไหนพร้อม
	fmt.Println("no value")
}
```

## รูปแบบที่ใช้บ่อย

- **Timeout:** case ที่รอ `time.After` หรือ `ctx.Done()`
- **ส่ง/รับแบบไม่บล็อก:** เพิ่ม `default`
- **Event loop:** `for { select { ... } }` ที่มี case สำหรับออก

:::caution[ข้อควรระวัง]
- `time.After` ใน loop ที่วนบ่อยสร้าง timer ใหม่ทุกรอบ ให้ใช้ `time.Timer` ซ้ำหรือใช้ context
- `select {}` เปล่าๆ บล็อกตลอดไป
:::
