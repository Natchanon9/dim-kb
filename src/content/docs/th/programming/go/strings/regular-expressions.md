---
title: "Regular Expressions"
description: "match ดึงข้อมูล และแทนที่ด้วย regexp"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

`regexp` ใช้ syntax แบบ RE2 ซึ่งรับประกันว่า match ในเวลาเชิงเส้น compile pattern ครั้งเดียวแล้วใช้ซ้ำ

## ตัวอย่าง

```go title="regexp.go"
var emailRe = regexp.MustCompile(`(\w+)@(\w+)\.com`) // compile ครั้งเดียว

emailRe.MatchString("a@b.com")                         // true
emailRe.FindStringSubmatch("mail bob@example.com now") // [bob@example.com bob example]
emailRe.ReplaceAllString("x@y.com", "$1 at $2")        // x at y
```

| Method | คืนค่า |
| --- | --- |
| `MatchString` | ตรงหรือไม่ |
| `FindString` / `FindAllString(s, -1)` | ตัวแรก / ทุกตัวที่ตรง |
| `FindStringSubmatch` | ตัวที่ตรงพร้อม capture group |
| `ReplaceAllString` | ข้อความที่แทนแล้ว (`$1` อ้างถึง group) |

ใช้ raw string (backtick) กับ pattern จะได้ไม่ต้อง escape `\w`

:::caution[ข้อควรระวัง]
- RE2 ไม่มี backreference และ lookahead
- `MustCompile` panic ถ้า pattern ผิด ใช้กับ pattern คงที่ระดับ package ได้ แต่กับ input จากผู้ใช้ให้ใช้ `Compile`
- เช็คง่ายๆ ใช้ `strings.Contains`/`HasPrefix` เร็วกว่าและอ่านง่ายกว่า
:::
