---
title: "HTTP Client"
description: "ส่ง request ออกไปพร้อม timeout และ context"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

สร้าง `http.Client` ที่มี timeout ตัวเดียวแล้วใช้ซ้ำ มันเก็บ connection ไว้ใช้ต่อ และใช้พร้อมกันหลาย goroutine ได้ปลอดภัย

## ตัวอย่าง

```go title="client.go"
client := &http.Client{Timeout: 5 * time.Second}

req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
req.Header.Set("Accept", "application/json")

res, err := client.Do(req)
if err != nil {
	return err // network error, timeout, ctx ถูกยกเลิก
}
defer res.Body.Close()
if res.StatusCode != http.StatusOK {
	return fmt.Errorf("get user: %s", res.Status) // 4xx/5xx ไม่ใช่ error จาก Do
}
var user User
err = json.NewDecoder(res.Body).Decode(&user)
```

## เช็คลิสต์

- มี timeout ที่ client และ/หรือ deadline ใน context
- `defer res.Body.Close()` เสมอ
- เช็ค `StatusCode` เอง
- ใช้ client ซ้ำ อย่าสร้างใหม่ทุก request

:::caution[ข้อควรระวัง]
- `http.Get` และ `http.DefaultClient` **ไม่มี timeout** server ที่ค้างจะทำให้บล็อกตลอดไป
- ไม่อ่าน/ไม่ปิด body ทำให้ใช้ connection ซ้ำไม่ได้
:::
