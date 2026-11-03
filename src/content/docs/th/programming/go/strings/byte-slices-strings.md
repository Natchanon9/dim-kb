---
title: "Byte Slice และ String"
description: "แปลงระหว่าง []byte กับ string และ package bytes"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

`[]byte` คือคู่ของ `string` ที่แก้ไขได้ งาน I/O, hashing และ JSON ใช้ `[]byte` ส่วน API ด้านข้อความส่วนใหญ่ใช้ `string`

## การแปลง

```go title="bytes.go"
b := []byte("hello") // copy byte ของ string มา
b[0] = 'H'           // []byte แก้ไขได้
fmt.Println(string(b), b, string(bytes.ToUpper(b)))
// Hello [72 101 108 108 111] HELLO
```

การแปลงแต่ละครั้ง copy ข้อมูล เพราะ string ต้องแก้ไขไม่ได้

## bytes เหมือน strings

`bytes` มี helper ชุดเดียวกับ `strings` (`Contains`, `Split`, `TrimSpace`, `Equal` …) และมี `bytes.Buffer` ซึ่งเป็น buffer ขยายได้ที่เป็นทั้ง `io.Reader` และ `io.Writer`

:::caution[ข้อควรระวัง]
- แปลงไปมาใน loop จะจอง memory ทุกครั้ง เลือก type เดียวแล้วใช้ไปตลอด
- เทียบ byte slice ด้วย `bytes.Equal` ไม่ใช่ `==` (ซึ่ง compile ไม่ผ่านกับ slice)
:::
