---
title: "Go Playground"
description: "รันและแชร์โค้ด Go บนเบราว์เซอร์"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

[go.dev/play](https://go.dev/play) compile และรัน Go ใน sandbox ใช้ลองไอเดียเล็กๆ หรือแชร์ snippet ที่คนอื่นรันซ้ำได้

## ทำอะไรได้บ้าง

- เลือกเวอร์ชัน Go ได้ (รวมถึง dev branch)
- **Share** ได้ลิงก์ถาวร เหมาะกับแจ้ง bug หรือถามคำถาม
- format โค้ดและรัน test ได้ ถ้าไม่มี `main` จะรัน function `Test` ให้
- หลายไฟล์ได้ โดยคั่นด้วยบรรทัด `-- name.go --`

## ข้อจำกัดของ sandbox

- ออก network ไม่ได้ และไม่มี filesystem จริง
- เวลาเริ่มที่ `2009-11-10 23:00:00 UTC` และ sleep เป็นแบบจำลอง output จึงออกมาเหมือนเดิมทุกครั้ง

:::caution[ข้อควรระวัง]
- อย่าวาง secret: snippet ที่แชร์แล้วใครมีลิงก์ก็เปิดได้
- `time.Now()` เป็นเวลาปลอม โค้ดที่ขึ้นกับเวลาจะทำงานต่างจากในเครื่อง
:::
