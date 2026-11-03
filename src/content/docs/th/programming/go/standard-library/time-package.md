---
title: "time Package"
description: "เวลา ระยะเวลา การ format และการ parse"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

`time.Time` คือจุดเวลาที่มี location ส่วน `time.Duration` คือ nanosecond ที่เป็น `int64` การ format ใช้ **วันอ้างอิง** แทน `%Y-%m-%d`

## ตัวอย่าง

```go title="time.go"
t := time.Date(2026, 11, 3, 14, 30, 0, 0, time.UTC)
t.Format("2006-01-02 15:04")            // 2026-11-03 14:30
t.Format(time.RFC3339)                  // 2026-11-03T14:30:00Z
t.Weekday()                             // Tuesday
t.Add(90 * time.Minute).Format(time.Kitchen) // 4:00PM

d, _ := time.ParseDuration("1h30m")     // d.Minutes() → 90
xmas, _ := time.Parse("02/01/2006", "25/12/2026")
xmas.Sub(t).Hours()                     // 1233.5
```

วันอ้างอิงคือ `Mon Jan 2 15:04:05 MST 2006` (1 2 3 4 5 6 7)

## นิสัยที่ดี

- เก็บและส่งเวลาเป็น UTC (RFC 3339) แปลงเป็นเวลาท้องถิ่นตอนแสดงผลเท่านั้น (`time.LoadLocation("Asia/Bangkok")`)
- เทียบด้วย `t.Before`, `t.After`, `t.Equal` ไม่ใช่ `==`
- วัดเวลาที่ใช้ไปด้วย `time.Since(start)`

:::caution[ข้อควรระวัง]
- `time.Parse` ที่ไม่มี zone จะได้ UTC ถ้า input เป็นเวลาท้องถิ่นให้ใช้ `ParseInLocation`
- `time.Sleep` ใน test ทำให้ test ช้าและ flaky
:::
