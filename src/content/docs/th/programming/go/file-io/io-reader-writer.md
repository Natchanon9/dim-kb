---
title: "io.Reader และ io.Writer"
description: "สอง interface ที่ API ด้าน I/O ทุกตัวของ Go สร้างขึ้นมา"
sidebar:
  order: 3
status: done
lastReviewed: 2026-11-03
---

`io.Reader` (`Read(p []byte) (n int, err error)`) และ `io.Writer` (`Write(p []byte) (n int, err error)`) เชื่อมไฟล์, network connection, HTTP body, ตัวบีบอัด และ hash เข้าด้วยกัน เขียนโค้ดกับสอง interface นี้แล้วใช้ได้กับทุกอย่างข้างต้น

## Writer ของเราเอง และ reader ที่ต่อกัน

```go title="io.go"
type upperWriter struct{ w io.Writer }

func (u upperWriter) Write(p []byte) (int, error) {
	return u.w.Write(bytes.ToUpper(p))
}

var sb strings.Builder
io.Copy(upperWriter{&sb}, strings.NewReader("go io")) // GO IO

r := io.LimitReader(strings.NewReader("abcdef"), 3)
b, _ := io.ReadAll(r) // abc
```

## Helper ที่มีประโยชน์

| Helper | ทำอะไร |
| --- | --- |
| `io.Copy(dst, src)` | stream ทุกอย่างจาก src ไป dst |
| `io.ReadAll(r)` | อ่านจนจบเข้า memory |
| `io.MultiWriter(a, b)` | เขียนลงหลาย writer |
| `io.TeeReader(r, w)` | copy สิ่งที่อ่านลง w ด้วย |
| `io.LimitReader(r, n)` | หยุดหลัง n byte |
| `io.Discard` | writer ที่ทิ้งทุกอย่าง |

:::caution[ข้อควรระวัง]
- `Read` อาจคืนน้อยกว่าที่ขอ และอาจคืนข้อมูลพร้อม `io.EOF` ถ้าต้องการให้ครบ n byte ใช้ `io.ReadFull`
:::
