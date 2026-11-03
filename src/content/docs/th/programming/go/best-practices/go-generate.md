---
title: "go generate"
description: "รันตัว generate โค้ดที่ประกาศไว้ใน comment"
sidebar:
  order: 7
status: done
lastReviewed: 2026-11-03
---

`go generate` หา comment `//go:generate` แล้วรันคำสั่งนั้น ไฟล์ที่ generate ได้จะถูก commit ไว้ build ปกติจึงไม่ต้องมีเครื่องมือเหล่านั้น

## ตัวอย่าง: stringer

```go title="color.go"
//go:generate stringer -type=Color

type Color int

const (
	Red Color = iota
	Green
	Blue
)
```

```sh
go install golang.org/x/tools/cmd/stringer@latest
go generate ./...   # เขียน color_string.go ที่มี func (Color) String()
```

```go
fmt.Println(Red, Blue, Color(7)) // Red Blue Color(7)
```

## ตัว generate ที่เจอบ่อย

| เครื่องมือ | สร้างอะไร |
| --- | --- |
| `stringer` | `String()` สำหรับ enum |
| `mockgen` | mock จาก interface |
| `sqlc` | โค้ด Go ที่มี type จาก SQL |
| `protoc-gen-go` | Protocol Buffers / gRPC |
| `swag` | Swagger doc จาก comment |

:::caution[ข้อควรระวัง]
- `go build` ไม่เคยรัน `go generate` ให้ รันใหม่ทุกครั้งที่แก้ source และให้ CI เช็คว่าไฟล์ที่ generate ยังเป็นปัจจุบัน
:::
