---
title: "Interface ที่ใช้บ่อย"
description: "fmt.Stringer, error, io.Reader, io.Writer และตัวอื่นๆ"
sidebar:
  order: 10
status: done
lastReviewed: 2026-11-03
---

มี interface มาตรฐานไม่กี่ตัวที่เจอทุกที่ implement พวกนี้แล้ว type ของเราจะต่อเข้ากับ standard library ได้ทันที

## Stringer และ io

```go title="common.go"
type Temp float64

func (t Temp) String() string { return fmt.Sprintf("%.1f°C", float64(t)) }

fmt.Println(Temp(21.5)) // 21.5°C

r := strings.NewReader("hello") // io.Reader
var sb strings.Builder          // io.Writer
n, _ := io.Copy(&sb, r)
fmt.Println(n, sb.String()) // 5 hello
```

| Interface | Method | ใช้โดย |
| --- | --- | --- |
| `fmt.Stringer` | `String() string` | การพิมพ์ของ `fmt` |
| `error` | `Error() string` | error handling ทั้งหมด |
| `io.Reader` | `Read(p []byte) (n int, err error)` | ไฟล์, HTTP body, gzip, JSON decoder |
| `io.Writer` | `Write(p []byte) (n int, err error)` | ไฟล์, response, buffer, hash |
| `io.Closer` | `Close() error` | อะไรก็ตามที่ถือ resource |
| `sort.Interface` | `Len`, `Less`, `Swap` | `sort.Sort` (ส่วนใหญ่ถูกแทนด้วย `slices.SortFunc` แล้ว) |

:::caution[ข้อควรระวัง]
- `Read` คืนข้อมูล **และ** `io.EOF` ในการเรียกครั้งเดียวกันได้ ให้จัดการ `n` byte ก่อนค่อยเช็ค error
:::
